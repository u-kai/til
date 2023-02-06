const asyncPromise = (i, k, j) =>
    new Promise((r) => {
        setTimeout(() => {
            console.log("start : ", k);
            for (let i = 0; i < j; i++) {}
            console.log("end : ", k);
            r("end");
        }, i);
    });

asyncPromise(1000, 1, 1);
asyncPromise(10, 2, 1);
asyncPromise(100000, 3, 1);
console.log("Hello Wolrd");
