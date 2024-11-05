// API Base URL
const API_URL = 'http://localhost:8080/api';

// Store data
let currentUser = null;

// Show/hide content sections
function showContent(contentId) {
    document.querySelectorAll('.content').forEach(content => {
        content.classList.add('hidden');
    });
    document.getElementById(contentId).classList.remove('hidden');

    if (contentId === 'user-list') {
        fetchUsers();
    }
}

// API Error Handler
async function handleApiResponse(response) {
    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.message || 'Something went wrong');
    }
    return response.json();
}

// Handle registration
async function handleRegistration(event) {
    event.preventDefault();
    const userData = {
        name: document.getElementById('name').value,
        username: document.getElementById('username').value,
        password: document.getElementById('password').value
    };

    try {
        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        });

        const data = await handleApiResponse(response);

        // Check if the token exists in the response
        if (data.Data && data.Data.token) {
            localStorage.setItem('token', data.Data.token);
            currentUser = data;
            event.target.reset();
            showContent('todo-list');
            fetchTodos();
        } else {
            alert("Registration successful, but no token received.");
        }
    } catch (error) {
        alert(error.message);
    }
}


// Todo functions
async function addTodo() {
    const title = document.getElementById('todoTitle').value;
    if (!title) return;

    try {
        const response = await fetch(`${API_URL}/todos`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: JSON.stringify({ title })
        });

        await handleApiResponse(response);
        document.getElementById('todoTitle').value = '';
        fetchTodos();
    } catch (error) {
        alert(error.message);
    }
}

async function toggleTodoStatus(id, currentStatus) {
    try {
        const response = await fetch(`${API_URL}/todos/${id}/status`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: JSON.stringify({
                status: currentStatus === 'progress' ? 'done' : 'progress'
            })
        });

        await handleApiResponse(response);
        fetchTodos();
    } catch (error) {
        alert(error.message);
    }
}

async function fetchTodos() {
    try {
        const response = await fetch(`${API_URL}/todos`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        const todos = await handleApiResponse(response);
        renderTodoList(todos);
    } catch (error) {
        alert(error.message);
    }
}

function renderTodoList(todos) {
    const todoItems = document.getElementById('todoItems');
    todoItems.innerHTML = todos.map(todo => `
        <div class="todo-item">
            <span>${todo.title}</span>
            <button 
                onclick="toggleTodoStatus('${todo.id}', '${todo.status}')" 
                class="status-btn ${todo.status === 'done' ? 'done' : ''}"
            >
                ${todo.status}
            </button>
        </div>
    `).join('');
}

// User list functions
async function fetchUsers() {
    try {
        const response = await fetch(`${API_URL}/users`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        const users = await handleApiResponse(response);
        renderUserList(users);
    } catch (error) {
        alert(error.message);
    }
}

function renderUserList(users) {
    const userItems = document.getElementById('userItems');
    userItems.innerHTML = users.map(user => `
        <div class="user-item">
            <span>${user.name}</span>
            <span class="status-indicator ${user.isActive ? 'status-active' : 'status-inactive'}">
                ${user.isActive ? 'active' : 'inactive'}
            </span>
            <button onclick="showUserDetail('${user.id}')" class="detail-btn">Detail</button>
        </div>
    `).join('');
}

async function showUserDetail(userId) {
    try {
        const response = await fetch(`${API_URL}/users/${userId}`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        const user = await handleApiResponse(response);
        renderUserDetail(user);
    } catch (error) {
        alert(error.message);
    }
}

function renderUserDetail(user) {
    const detailContent = document.getElementById('userDetailContent');
    detailContent.innerHTML = `
        <div class="form-group">
            <label>Name:</label>
            <input type="text" value="${user.name}" readonly>
        </div>
        <div class="form-group">
            <label>Username:</label>
            <input type="text" value="${user.username}" readonly>
        </div>
        <div class="form-group">
            <label>Password:</label>
            <input type="text" value="${user.password}" readonly>
        </div>
        <div class="form-group">
            <label>Token:</label>
            <input type="text" value="${user.token}" readonly>
        </div>
    `;
    showContent('user-detail');
}

// Check authentication status
function checkAuth() {
    const token = localStorage.getItem('token');
    if (!token) {
        showContent('registration');
    } else {
        // Verify token validity with backend
        fetch(`${API_URL}/verify-token`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            .then(handleApiResponse)
            .then(data => {
                currentUser = data;
                showContent('todo-list');
                fetchTodos();
            })
            .catch(() => {
                localStorage.removeItem('token');
                showContent('registration');
            });
    }
}

// Logout function
function logout() {
    localStorage.removeItem('token');
    currentUser = null;
    showContent('registration');
}

// Initialize app
document.addEventListener('DOMContentLoaded', checkAuth);



// response dari api

// {
//     "id": "string",
//     "name": "string",
//     "username": "string",
//     "password": "string",
//     "token": "string",
//     "isActive": boolean
// }

// {
//     "id": "string",
//     "title": "string",
//     "status": "string" // "progress" atau "done"
// }