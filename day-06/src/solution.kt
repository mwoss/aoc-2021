import java.io.File

fun getNumberOfFishesAfter(fishes: MutableList<Int>, days: Int): Int {
    for (`_` in 1..days) {
        val currentFishAmount = fishes.size - 1
        for (i in 0..currentFishAmount) {
            if (fishes[i] == 0) {
                fishes.add(8)
                fishes[i] = 6
            } else {
                fishes[i]--
            }
        }
    }
    return fishes.size
}

fun main() {
    val lanternFishes = File("src/input.txt")
        .readText()
        .trim()
        .split(",")
        .map { it.toInt() }
        .toMutableList()

    println(getNumberOfFishesAfter(lanternFishes, 80))
    println(getNumberOfFishesAfter(lanternFishes, 256))
}