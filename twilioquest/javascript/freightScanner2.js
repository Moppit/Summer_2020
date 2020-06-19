function scan(arr) {
    var appears = [];
    for(var i = 0; i < arr.length; i++) {
        if(arr[i] == "contraband") {
            appears.push(i);
        }
    }
    return appears;
}

// Second char: [3][4] 