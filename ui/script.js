let selectedFunc = '';
let selectedHash = '';

function selectFunc(func) {
    selectedFunc = func;
    document.getElementById('encrypt').classList.remove('active');
    document.getElementById('decrypt').classList.remove('active');
    document.getElementById(func).classList.add('active');
}

function selectHash(hash) {
    selectedHash = hash;
    document.getElementById('sha256').classList.remove('active');
    document.getElementById('sha384').classList.remove('active');
    document.getElementById('sha512').classList.remove('active');
    document.getElementById(hash).classList.add('active');
}

function startProcess() {
    const inputText = document.getElementById('input').value;
    if (!selectedFunc || !selectedHash || !inputText) {
        alert('Please select function, hash type, and input text.');
        return;
    }

    fetch('/api/' + selectedFunc, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            text: inputText,
            hash: selectedHash
        })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('output').value = data.result;
        addHistory(inputText, data.result);
    })
    .catch(error => console.error('Error:', error));
}

function addHistory(key, crypto) {
    const historyBody = document.getElementById('history-body');
    const row = historyBody.insertRow();
    const cellKey = row.insertCell(0);
    const cellCrypto = row.insertCell(1);

    cellKey.textContent = key;
    cellCrypto.textContent = crypto;
}