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
	try {
		const f = fibo();
		console.log(f);
	} catch {
		console.log("on no error");
	}
	for (let i = 0; i < 10; i++) {
		console.log(f());
	}
}

main();
let bar = new Foo();
const asdf = 69;

class Foo {
	print = function () {
		console.log(this.value);
	};
}

{
	{
		{
			{
				{
					{
					}
				}
			}
		}
	}
}
