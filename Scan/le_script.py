import requests
import json
import re
from bs4 import BeautifulSoup
from flask import Flask, render_template

WP_API_URL = "https://wpscan.com/api/v3/wordpresses/"

app = Flask(__name__)

def get_wordpress_version(url):
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        meta_generator = soup.find('meta', attrs={'name': 'generator'})
        if meta_generator and 'WordPress' in meta_generator['content']:
            return meta_generator['content'].split(' ')[1]
    except Exception as e:
        print(f"Erreur lors de la récupération de la version WordPress: {e}")
    return None

def get_server_info(url):
    try:
        response = requests.get(url)
        server = response.headers.get("Server", "Inconnu")
        php_version = response.headers.get("X-Powered-By", "Inconnu")
        return {"server": server, "php_version": php_version}
    except Exception as e:
        print(f"Erreur lors de la récupération des informations serveur: {e}")
    return {}

def check_vulnerabilities(version):
    if not version:
        return "Impossible de déterminer la version de WordPress."
    try:
        version_number = re.sub("\.", "", version)
        print(version_number)
        api_url = f"{WP_API_URL}{version_number}"
        print(api_url)
        response = requests.get(api_url, headers = { "Authorization: Token", "token=1l2Z5A2Vt1nbcQlPaA8hui03AaRjncYLae3RGFzlhb4" })
        print(response.json())
        return response.json()
    except Exception as e:
        return {"error": f"Erreur API: {e}"}

def audit_site(url):
    version = get_wordpress_version(url)
    server_info = get_server_info(url)
    vulnerabilities = check_vulnerabilities(version)
    report = {
        "Version WordPress": version,
        "Serveur": server_info
    }
    # Vulnérabilités: {json.dumps(vulnerabilities, indent=2)}
    return report

@app.route("/")
def index():
    url = input("url du site:")
    report = audit_site(url)
    return json.dumps(report)

if __name__ == "__main__":
    app.run(debug=True)