// DOM Elements
const app = document.getElementById('app');
const darkModeToggle = document.getElementById('darkModeToggle');
const tabsList = document.querySelector('.tabs-list');
const tabContent = document.getElementById('tabContent');

// State
let darkMode = true;
let activeTab = 'login';
let users = [];
let todos = [];
let selectedUser = null;
let loggedInUser = null;

// API Base URL
const API_BASE_URL = 'http://localhost:8080/api/'; // Replace with your actual API base URL

// Fetch and render tab content
async function renderTabContent(tab) {
    const response = await fetch(`templates/${tab}.html`);
    const html = await response.text();
    tabContent.innerHTML = html;

    if (tab === 'todo-list') {
        renderTodos();
    } else if (tab === 'user-list') {
        renderUsers();
    } else if (tab === 'user-detail' && selectedUser) {
        renderUserDetail();
    }
}

// Toggle dark mode
function toggleDarkMode() {
    darkMode = !darkMode;
    app.classList.toggle('dark', darkMode);
    darkModeToggle.innerHTML = darkMode ?
        '<i data-lucide="sun" class="h-[1.2rem] w-[1.2rem]"></i>' :
        '<i data-lucide="moon" class="h-[1.2rem] w-[1.2rem]"></i>';
    lucide.createIcons();
}

// Set active tab
function setActiveTab(tab) {
    activeTab = tab;
    document.querySelectorAll('.tab-trigger').forEach(trigger => {
        trigger.classList.toggle('active', trigger.dataset.tab === tab);
    });
    renderTabContent(tab);
}

// Fetch users from API
async function fetchUsers() {
    try {
        const response = await fetch(`${API_BASE_URL}/users`);
        users = await response.json();
        renderUsers();
    } catch (error) {
        console.error('Error fetching users:', error);
    }
}

// Fetch todos from API
async function fetchTodos() {
    try {
        const response = await fetch(`${API_BASE_URL}/todos`);
        todos = await response.json();
        renderTodos();
    } catch (error) {
        console.error('Error fetching todos:', error);
    }
}

// Render users
function renderUsers() {
    const userList = document.getElementById('userList');
    if (!userList) return;

    userList.innerHTML = users.map(user => `
        <li class="${user.active ? 'bg-green-100 dark:bg-green-900' : 'bg-red-100 dark:bg-red-900'}">
            <span class="${user.active ? 'text-green-800 dark:text-green-200' : 'text-red-800 dark:text-red-200'}">${user.name}</span>
            <div>
                <span class="status-badge ${user.active ? 'bg-green-200 text-green-800' : 'bg-red-200 text-red-800'}">
                    ${user.active ? 'Active' : 'Inactive'}
                </span>
                <button onclick="showUserDetail('${user.id}')" class="btn btn-outline btn-sm">Detail</button>
                <button onclick="deleteUser('${user.id}')" class="btn btn-danger btn-icon">
                    <i data-lucide="trash-2" class="h-4 w-4"></i>
                </button>
            </div>
        </li>
    `).join('');
    lucide.createIcons();
}

// Render todos
function renderTodos() {
    const todoList = document.getElementById('todoList');
    if (!todoList) return;

    todoList.innerHTML = todos.map(todo => `
        <li>
            <div class="flex items-center space-x-2">
                <input type="checkbox" id="todo-${todo.id}" ${todo.done ? 'checked' : ''} onchange="toggleTodoStatus('${todo.id}')">
                <label for="todo-${todo.id}" class="${todo.done ? 'line-through text-muted' : ''}">${todo.title}</label>
            </div>
            <button onclick="deleteTodo('${todo.id}')" class="btn btn-danger btn-icon">
                <i data-lucide="trash-2" class="h-4 w-4"></i>
            </button>
        </li>
    `).join('');
    lucide.createIcons();
}

// Render user detail
function renderUserDetail() {
    if (!selectedUser) return;
    document.getElementById('userName').textContent = selectedUser.name;
    document.getElementById('userUsername').textContent = selectedUser.username;
    document.getElementById('userPassword').textContent = selectedUser.password;
    document.getElementById('userToken').textContent = selectedUser.token;
}

// Show user detail
async function showUserDetail(userId) {
    try {
        const response = await fetch(`${API_BASE_URL}/users/${userId}`);
        selectedUser = await response.json();
        setActiveTab('user-detail');
    } catch (error) {
        console.error('Error fetching user detail:', error);
    }
}

// Register user
async function registerUser(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const userData = Object.fromEntries(formData);

    try {
        const response = await fetch(`${API_BASE_URL}/auth/register`, {

            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(userData)
        });
        const newUser = await response.json();
        users.push(newUser);
        setActiveTab('login');
    } catch (error) {
        console.error('Error registering user:', error);
    }
}

// Login user
async function loginUser(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const { username, password } = Object.fromEntries(formData);

    try {
        const response = await fetch(`${API_BASE_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        });
        loggedInUser = await response.json();
        setActiveTab('todo-list');
        fetchTodos();
    } catch (error) {
        console.error('Error logging in:', error);
        alert('Invalid credentials');
    }
}

// Add todo
async function addTodo(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const todoData = Object.fromEntries(formData);

    try {
        const response = await fetch(`${API_BASE_URL}/todos`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(todoData)
        });
        const newTodo = await response.json();
        todos.push(newTodo);
        renderTodos();
        event.target.reset();
    } catch (error) {
        console.error('Error adding todo:', error);
    }
}

// Toggle todo status
async function toggleTodoStatus(todoId) {
    try {
        const response = await fetch(`${API_BASE_URL}/todos/${todoId}/toggle`, {
            method: 'PUT'
        });
        const updatedTodo = await response.json();
        todos = todos.map(todo => todo.id === updatedTodo.id ? updatedTodo : todo);
        renderTodos();
    } catch (error) {
        console.error('Error toggling todo status:', error);
    }
}

// Delete todo
async function deleteTodo(todoId) {
    try {
        await fetch(`${API_BASE_URL}/todos/${todoId}`, {
            method: 'DELETE'
        });
        todos = todos.filter(todo => todo.id !== todoId);
        renderTodos();
    } catch (error) {
        console.error('Error deleting todo:', error);
    }
}

// Delete user
async function deleteUser(userId) {
    try {
        await fetch(`${API_BASE_URL}/users/${userId}`, {
            method: 'DELETE'
        });
        users = users.filter(user => user.id !== userId);
        if (selectedUser && selectedUser.id === userId) {
            selectedUser = null;
            setActiveTab('user-list');
        } else {
            renderUsers();
        }
    } catch (error) {
        console.error('Error deleting user:', error);
    }
}

// Event listeners
darkModeToggle.addEventListener('click', toggleDarkMode);
tabsList.addEventListener('click', (e) => {
    if (e.target.classList.contains('tab-trigger')) {
        setActiveTab(e.target.dataset.tab);
    }
});

// Initialize app
toggleDarkMode(); // Set initial dark mode
setActiveTab('login');
fetchUsers();

// Add event listeners for forms
document.addEventListener('submit', (e) => {
    if (e.target.id === 'loginForm') {
        loginUser(e);
    } else if (e.target.id === 'registerForm') {
        registerUser(e);
    } else if (e.target.id === 'addTodoForm') {
        addTodo(e);
    }
});

// Initialize Lucide icons
lucide.createIcons();


// // API Base URL
// const API_URL = 'http://localhost:8080/api';

// // Store data
// let currentUser = null;

// // Show/hide content sections
// function showContent(contentId) {
//     document.querySelectorAll('.content').forEach(content => {
//         content.classList.add('hidden');
//     });
//     document.getElementById(contentId).classList.remove('hidden');

//     if (contentId === 'user-list') {
//         fetchUsers();
//     }
// }

// // API Error Handler
// async function handleApiResponse(response) {
//     if (!response.ok) {
//         const error = await response.json();
//         throw new Error(error.message || 'Something went wrong');
//     }
//     return response.json();
// }

// // Handle registration
// async function handleRegistration(event) {
//     event.preventDefault();
//     const userData = {
//         name: document.getElementById('name').value,
//         username: document.getElementById('username').value,
//         password: document.getElementById('password').value
//     };

//     try {
//         const response = await fetch(`${API_URL}/auth/register`, {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify(userData)
//         });

//         const data = await handleApiResponse(response);

//         // Check if the token exists in the response
//         if (data.Data && data.Data.token) {
//             localStorage.setItem('token', data.Data.token);
//             currentUser = data;
//             event.target.reset();
//             showContent('todo-list');
//             fetchTodos();
//         } else {
//             alert("Registration successful, but no token received.");
//         }
//     } catch (error) {
//         alert(error.message);
//     }
// }


// // Todo functions
// async function addTodo() {
//     const title = document.getElementById('todoTitle').value;
//     if (!title) return;

//     try {
//         const response = await fetch(`${API_URL}/todos`, {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json',
//                 'Authorization': `Bearer ${localStorage.getItem('token')}`
//             },
//             body: JSON.stringify({ title })
//         });

//         await handleApiResponse(response);
//         document.getElementById('todoTitle').value = '';
//         fetchTodos();
//     } catch (error) {
//         alert(error.message);
//     }
// }

// async function toggleTodoStatus(id, currentStatus) {
//     try {
//         const response = await fetch(`${API_URL}/todos/${id}/status`, {
//             method: 'PUT',
//             headers: {
//                 'Content-Type': 'application/json',
//                 'Authorization': `Bearer ${localStorage.getItem('token')}`
//             },
//             body: JSON.stringify({
//                 status: currentStatus === 'progress' ? 'done' : 'progress'
//             })
//         });

//         await handleApiResponse(response);
//         fetchTodos();
//     } catch (error) {
//         alert(error.message);
//     }
// }

// async function fetchTodos() {
//     try {
//         const response = await fetch(`${API_URL}/todos`, {
//             headers: {
//                 'Authorization': `Bearer ${localStorage.getItem('token')}`
//             }
//         });

//         const todos = await handleApiResponse(response);
//         renderTodoList(todos);
//     } catch (error) {
//         alert(error.message);
//     }
// }

// function renderTodoList(todos) {
//     const todoItems = document.getElementById('todoItems');
//     todoItems.innerHTML = todos.map(todo => `
//         <div class="todo-item">
//             <span>${todo.title}</span>
//             <button 
//                 onclick="toggleTodoStatus('${todo.id}', '${todo.status}')" 
//                 class="status-btn ${todo.status === 'done' ? 'done' : ''}"
//             >
//                 ${todo.status}
//             </button>
//         </div>
//     `).join('');
// }

// // User list functions
// async function fetchUsers() {
//     try {
//         const response = await fetch(`${API_URL}/users`, {
//             headers: {
//                 'Authorization': `Bearer ${localStorage.getItem('token')}`
//             }
//         });

//         const users = await handleApiResponse(response);
//         renderUserList(users);
//     } catch (error) {
//         alert(error.message);
//     }
// }

// function renderUserList(users) {
//     const userItems = document.getElementById('userItems');
//     userItems.innerHTML = users.map(user => `
//         <div class="user-item">
//             <span>${user.name}</span>
//             <span class="status-indicator ${user.isActive ? 'status-active' : 'status-inactive'}">
//                 ${user.isActive ? 'active' : 'inactive'}
//             </span>
//             <button onclick="showUserDetail('${user.id}')" class="detail-btn">Detail</button>
//         </div>
//     `).join('');
// }

// async function showUserDetail(userId) {
//     try {
//         const response = await fetch(`${API_URL}/users/${userId}`, {
//             headers: {
//                 'Authorization': `Bearer ${localStorage.getItem('token')}`
//             }
//         });

//         const user = await handleApiResponse(response);
//         renderUserDetail(user);
//     } catch (error) {
//         alert(error.message);
//     }
// }

// function renderUserDetail(user) {
//     const detailContent = document.getElementById('userDetailContent');
//     detailContent.innerHTML = `
//         <div class="form-group">
//             <label>Name:</label>
//             <input type="text" value="${user.name}" readonly>
//         </div>
//         <div class="form-group">
//             <label>Username:</label>
//             <input type="text" value="${user.username}" readonly>
//         </div>
//         <div class="form-group">
//             <label>Password:</label>
//             <input type="text" value="${user.password}" readonly>
//         </div>
//         <div class="form-group">
//             <label>Token:</label>
//             <input type="text" value="${user.token}" readonly>
//         </div>
//     `;
//     showContent('user-detail');
// }

// // Check authentication status
// function checkAuth() {
//     const token = localStorage.getItem('token');
//     if (!token) {
//         showContent('registration');
//     } else {
//         // Verify token validity with backend
//         fetch(`${API_URL}/verify-token`, {
//                 headers: {
//                     'Authorization': `Bearer ${token}`
//                 }
//             })
//             .then(handleApiResponse)
//             .then(data => {
//                 currentUser = data;
//                 showContent('todo-list');
//                 fetchTodos();
//             })
//             .catch(() => {
//                 localStorage.removeItem('token');
//                 showContent('registration');
//             });
//     }
// }

// // Logout function
// function logout() {
//     localStorage.removeItem('token');
//     currentUser = null;
//     showContent('registration');
// }

// // Initialize app
// document.addEventListener('DOMContentLoaded', checkAuth);



// // response dari api

// // {
// //     "id": "string",
// //     "name": "string",
// //     "username": "string",
// //     "password": "string",
// //     "token": "string",
// //     "isActive": boolean
// // }

// // {
// //     "id": "string",
// //     "title": "string",
// //     "status": "string" // "progress" atau "done"
// // }