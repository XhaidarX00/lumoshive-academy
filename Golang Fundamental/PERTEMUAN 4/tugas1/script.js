// contoh implementasi ajax

// const req = new XMLHttpRequest;

// const baseUrl = 'www.google.com'

// req.open('GET', baseUrl, '/users')

// req.onreadystatechange = function() {
//     if (req.status == 200) {
//         var data = JSON.parse(req.responseText())
//         console.log(data.name);
//     }
// }

// req.send()


async function fetchGameList(page = 1) {
    try {
        const response = await fetch(`https://lumoshive-academy-media-api.vercel.app/api/games?page=${page}&search`);
        const games = await response.json();

        const gameListContainer = document.getElementById('game-list');
        gameListContainer.innerHTML = '';

        console.log(games);
        games.forEach(game => {
            const gameCard = document.createElement('div');
            gameCard.className = 'game-card';

            gameCard.innerHTML = `
                <img src="${game.thumb}" alt="${game.title}" class="game-thumbnail">
                <div class="game-details">
                    <h2 class="game-title">${game.title}</h2>
                    <p class="game-description">${game.desc}</p>
                </div>
            `;

            gameListContainer.appendChild(gameCard);
        });
    } catch (error) {
        console.error('Error fetching the game list:', error);
    }
}

document.getElementById('page-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const pageInput = document.getElementById('page-input');
    const pageNumber = pageInput.value;
    fetchGameList(pageNumber);
});

document.addEventListener('DOMContentLoaded', () => fetchGameList(1));