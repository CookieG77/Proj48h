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