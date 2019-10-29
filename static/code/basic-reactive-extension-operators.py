import rx
from rx import operators as ops
import random

# Map
source = rx.of(0, 1, 2, 3, 4, 5)
source.pipe(
    ops.map(lambda number: number**2)
).subscribe(
    lambda number: print(number),
    lambda error: print(error),
    lambda: print("Done\n")
)

# Filter
source = rx.of(0, 1, 2, 3, 4, 5)
source.pipe(
    ops.filter(lambda number: number % 2 == 0)
).subscribe(
    lambda number: print(number),
    lambda error: print(error),
    lambda: print("Done\n")
)

# Flatmap
book_ids = range(0, 10)
library = [(1, "Dracula"), (2, "Harry Potter"), (3, "The Divine Comedy")]

def fetch_book_by_id(book_id):
    return rx.from_iterable(library).pipe(
        ops.filter(lambda book_record: book_record[0] == book_id),
        ops.map(lambda book_record: book_record[1])
    )


source = rx.from_iterable(book_ids)
source.pipe(
    ops.flat_map(lambda book_id: fetch_book_by_id(book_id))
).subscribe(
    lambda number: print(number),
    lambda error: print(error),
    lambda: print("Done\n")
)

# First
source = rx.of(0, 1, 2, 3, 4, 5)
source.pipe(
    ops.first()
).subscribe(
    lambda number: print(number),
    lambda error: print(error),
    lambda: print("Done\n")
)

# last
source = rx.of(0, 1, 2, 3, 4, 5)
source.pipe(
    ops.last(lambda number: number % 2 == 0)
).subscribe(
    lambda number: print(number),
    lambda error: print(error),
    lambda: print("Done\n")
)

# Buffer

binary_input_str = "01000101011110000110000101101101011100000110110001100101"
binary_digits = [i for i in binary_input_str]

source = rx.from_iterable(binary_digits).pipe(
    ops.buffer_with_count(8),
    ops.map(lambda bits_of_a_byte: "".join(bits_of_a_byte))
).subscribe(
    lambda byte: print(byte),
    lambda error: print(error),
    lambda: print("Done\n")
)

# take and take last

def take_first(haystack, n):
    return rx.from_iterable(haystack).pipe(
        ops.take(n),
        ops.buffer_with_count(n)
    )

def take_last(haystack, n):
    return rx.from_iterable(haystack).pipe(
        ops.take_last(n),
        ops.buffer_with_count(n)
    )

def find_needle(haystack, needle):
    middle = haystack[len(haystack) // 2]
    if needle == middle:
        return rx.of(middle)
    elif needle < middle:
        return take_first(haystack, len(haystack) // 2).pipe(
            ops.do_action(lambda haystack: print("Haystack size is ", len(haystack))),
            ops.flat_map(lambda haystack: find_needle(haystack, needle)))
    else:
        return take_last(haystack, len(haystack) // 2).pipe(
            ops.do_action(lambda haystack: print("Haystack size is ", len(haystack))),
            ops.flat_map(lambda haystack: find_needle(haystack, needle)))


haystack = [i for i in range(0, 1000)]
needle = random.choice(haystack)
print("Haystack size is ", len(haystack))
rx.of(haystack).pipe(
    ops.flat_map(lambda haystack: find_needle(haystack, needle))
).subscribe(
    lambda found_needle: print(f"\nThe needle found is {found_needle}. The original needle was {needle}"),
    lambda error: print(error),
    lambda: print("Done\n")
)

# Distinct

source = rx.of("cat", "Cat", "CAT", "dog", "dOg", "DOG")
source.pipe(
    ops.distinct(lambda x: x, lambda a, b: str.lower(a) == str.lower(b))
).subscribe(
    lambda number: print(number),
    lambda error: print(error),
    lambda: print("Done\n")
)
