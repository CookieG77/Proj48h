console.log("Hello from reportScript.js");

// Add event listener to download button to download the file when clicked on using AJAX
document.getElementById('downloadButton').addEventListener('click', function() {
    // Faire une requête AJAX pour récupérer le fichier
    fetch('/download')
        .then(response => response.blob()) // Récupère la réponse sous forme de blob
        .then(blob => {
            // Créer un lien de téléchargement dynamique
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.style.display = 'none';
            a.href = url;
            a.download = 'report.pdf'; // Spécifie le nom du fichier à télécharger
            document.body.appendChild(a);
            a.click();
            window.URL.revokeObjectURL(url); // Libère l'URL après le téléchargement
        })
        .catch(error => console.error('Error downloading file:', error));
});


// Add event listener to submit email button to send the email when clicked on using AJAX
document.getElementById('sendByMail').addEventListener('click', function() {
    const email = document.getElementById('email-input').value;

    // Check if email is valid
    if (!email) {
        alert("Please enter a valid email.");
        return;
    }

    // Use AJAX to send email with a POST request
    fetch('/send-mail-for-download', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: email }),
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Email submitted successfully!');
            } else {
                alert('Error submitting email.');
            }
        })
        .catch(error => console.error('Error submitting email:', error));
});