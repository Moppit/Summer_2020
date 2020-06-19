class TargetingSolution {
    constructor(obj) {
        this.x = obj.x;
        this.y = obj.y;
        this.z = obj.z;
    }

    target() {
        return "(" + this.x + ", " + this.y + ", " + this.z + ")";
    }   
}

// var testing = new TargetingSolution({
//     x: 45,
//     y: 12,
//     z: -1
//   });
// console.log(testing.target());