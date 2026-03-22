def greet(name: str) -> str:
    return f"Hello, {name}!"

def calculate_area(width: float, height: float) -> float:
    return width * height

def process_data(data: list) -> list:
    return [x ** 2 for x in data]

if __name__ == "__main__":
    print(greet("John"))
    print(calculate_area(10, 5))
    print(process_data([1, 2, 3, 4, 5]))