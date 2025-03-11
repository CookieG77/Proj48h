 function handleSubmit() {
    const urlInput = document.getElementById('websiteUrl');
    const statusMessage = document.getElementById('statusMessage');

    if (!urlInput.checkValidity()) {
    urlInput.reportValidity();
    return;
}

    statusMessage.textContent = "Analyzing website...";
    statusMessage.style.opacity = '1';

    setTimeout(() => {
    statusMessage.textContent = "Analysis complete!";
    setTimeout(() => {
    statusMessage.style.opacity = '0';
}, 2000);
}, 1500);
}

    document.getElementById('websiteUrl')
    .addEventListener('keypress', function(e) {
    if (e.key === 'Enter') handleSubmit();
});

 document.addEventListener('DOMContentLoaded', () => {
     // Get all typewriter elements
     const typewriters = document.querySelectorAll('.typewriter');

     typewriters.forEach(element => {
         // Get text content and calculate length
         const text = element.textContent.trim();
         const length = text.length;

         // Set CSS custom property
         element.style.setProperty('--char-count', length);

     });
 });

 function setLang(lang) {
     document.cookie = `lang=${lang};expires=0;samesite=lax;path=/`;
 }

 document.addEventListener('DOMContentLoaded', function () {
     const langSelect = document.querySelector('.lang-select');
     if (langSelect) {
         langSelect.addEventListener('change', function () {
             const selectedLang = this.value;
             setLang(selectedLang);
             console.log(selectedLang)
             location.reload();
         });
     }
 });

 function setTheme(theme) {
     document.cookie = `theme=${theme};expires=0;samesite=lax;path=/`;
 }