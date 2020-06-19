function scan(arr) {
    var appears = 0;
    arr.forEach(element => {
        if(element == "contraband") {
            appears++;
        }
    });
    return appears;
}

// Found first character: [1][3]