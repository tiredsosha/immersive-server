function sender(theme) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "http://127.0.0.1/themes");
    xhr.setRequestHeader("Accept", "application/json");
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            console.log(xhr.status);
            console.log(xhr.responseText);
        }
    };


    let data = `{
    "theme": ${theme}
    }`;

    xhr.send(data)
}