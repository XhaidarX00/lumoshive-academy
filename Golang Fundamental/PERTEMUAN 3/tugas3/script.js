document.addEventListener('DOMContentLoaded', function() {
    const dataForm = document.getElementById('dataForm');
    const dataInput = document.getElementById('dataInput');
    const dataList = document.getElementById('dataList');

    let dataStorage = JSON.parse(localStorage.getItem('dataList')) || [];

    if (dataStorage.length > 0) {
        renderListData();
    }

    dataForm.addEventListener('submit', function(event) {
        event.preventDefault();
        const data = dataInput.value.trim();

        if (data) {
            dataStorage.push(data);
            localStorage.setItem('dataList', JSON.stringify(dataStorage));
            dataInput.value = '';
            alert('Data berhasil disimpan!');
            renderListData();
        }
    });

    function renderListData() {
        dataList.innerHTML = '';
        dataStorage.forEach((item, index) => {
            const li = document.createElement('li');
            li.textContent = item;

            const deleteButton = document.createElement('button');
            deleteButton.textContent = 'Hapus';
            deleteButton.style.marginLeft = '10px';
            deleteButton.style.backgroundColor = '#a52323';
            deleteButton.style.color = 'white';
            deleteButton.style.border = 'none';
            deleteButton.style.listStyle = 'none';
            deleteButton.style.padding = '5px 10px';
            deleteButton.style.borderRadius = '5px';
            deleteButton.style.cursor = 'pointer';

            deleteButton.addEventListener('click', function() {
                dataStorage.splice(index, 1);
                localStorage.setItem('dataList', JSON.stringify(dataStorage));
                renderListData();
            });

            li.appendChild(deleteButton);
            dataList.appendChild(li);
        });
    }


    window.showPage = function(page) {
        const inputPage = document.getElementById('inputPage');
        const listPage = document.getElementById('listPage');

        if (page === 'input') {
            inputPage.style.display = 'block';
            listPage.style.display = 'none';
        } else if (page === 'list') {
            inputPage.style.display = 'none';
            listPage.style.display = 'block';
            renderListData();
        }
    };


    showPage('input');
});