async function fetchUserList(page = 1) {
    try {
        const response = await fetch(`https://reqres.in/api/users?page=${page}`);
        const data = await response.json();
        const users = data.data;

        const userListContainer = document.getElementById('user-list');
        userListContainer.innerHTML = '';

        users.forEach(user => {
            const userCard = document.createElement('div');
            userCard.className = 'user-card';
            userCard.dataset.userId = user.id;

            userCard.innerHTML = `
                <img src="${user.avatar}" alt="${user.first_name} ${user.last_name}" class="user-avatar">
                <div class="user-details">
                    <h2 class="user-name">${user.first_name} ${user.last_name}</h2>
                    <p class="user-email">${user.email}</p>
                </div>
            `;

            userCard.addEventListener('click', () => fetchUserDetails(user.id));
            userListContainer.appendChild(userCard);
        });
    } catch (error) {
        console.error('Error fetching the user list:', error);
    }
}

async function fetchUserDetails(userId) {
    try {
        const response = await fetch(`https://reqres.in/api/users/${userId}`);
        const data = await response.json();
        const user = data.data;

        const userDetailsContainer = document.getElementById('user-details');
        userDetailsContainer.innerHTML = `
            <div style="display: flex; flex-direction: column; align-items: center;">
                <img src="${user.avatar}" alt="${user.first_name} ${user.last_name}" style="width: 100px; height: 100px; border-radius: 50%; margin-bottom: 30px;">
                <div style="text-align: left; width: 120%; max-width: 300px;">
                    <h2 style="margin: 2px; margin-left: 30px;">${user.first_name} ${user.last_name}</h2>
                    <p style="margin: 2px; margin-left: 30px;">Email   : ${user.email}</p>
                    <p style="margin: 2px; margin-left: 30px;">User ID : ${user.id}</p>
                </div>
                <br>
            </div>
        `;

        showModal();
    } catch (error) {
        console.error('Error fetching user details:', error);
    }
}

function showModal() {
    const modal = document.getElementById('user-modal');
    modal.style.display = 'block';
}

function closeModal() {
    const modal = document.getElementById('user-modal');
    modal.style.display = 'none';
}

document.getElementById('page-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const pageInput = document.getElementById('page-input');
    const pageNumber = pageInput.value;
    fetchUserList(pageNumber);
});

document.querySelector('.close').addEventListener('click', closeModal);

window.addEventListener('click', function(event) {
    const modal = document.getElementById('user-modal');
    if (event.target === modal) {
        closeModal();
    }
});

document.addEventListener('DOMContentLoaded', () => fetchUserList(1));