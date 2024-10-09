document.getElementById('submit-btn').addEventListener('click', addText);

function addText() {
    const input = document.getElementById('text-input');
    const text = input.value.trim();
    if (text) {
        const listItem = document.createElement('div');
        listItem.className = 'text-item';
        listItem.innerHTML = `
            <span>${text}</span>
            <button class="delete-btn" onclick="this.parentElement.remove()">X</button>
        `;
        document.getElementById('text-list').appendChild(listItem);
        input.value = '';
    }
}