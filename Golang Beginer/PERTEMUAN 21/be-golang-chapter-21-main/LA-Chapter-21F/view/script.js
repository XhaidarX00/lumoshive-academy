document.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector('form');

    form.addEventListener('submit', function(event) {
        event.preventDefault();

        const username = document.getElementById('username').value.trim();
        const password = document.getElementById('password').value.trim();

        if (username === '' || password === '') {
            alert('Username and password are required!');
            return;
        }

        // Simulate form submission (replace with actual submission logic)
        alert(`Submitting: Username - ${username}, Password - ${password}`);
        form.reset(); // Reset form after submission
    });
});
