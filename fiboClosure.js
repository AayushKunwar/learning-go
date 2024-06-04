function fibo() {
	// let curr = 0,
	// prev = 1;
	let [prev, curr] = [1, 0];
	return () => {
		let next = curr + prev;
		prev = curr;
		curr = next;
		return prev;
	};
}

function main() {
	let f = fibo();
	for (let i = 0; i < 10; i++) {
		console.log(f());
	}
}

main();
