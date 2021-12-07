import java.io.File

const val BIRTH_DELAY = 7
const val BIRTH_DELAY_NEW_FISH = 9

fun getNumberOfFishesAfter(fishes: List<Int>, days: Int): Long {
    val birthDates = mutableMapOf<Int, Long>().apply {
        for (fish in fishes) {
            // incremented by 1 as new fishes are born on 0th day
            // note: merge works similar to default dict in Python :3
            merge(fish + 1, 1L, Long::plus)
        }
    }

    var population = fishes.size.toLong()
    for (day in 0..days) {
        birthDates[day]?.let {
            population += it
            birthDates.merge(day + BIRTH_DELAY, it, Long::plus)
            birthDates.merge(day + BIRTH_DELAY_NEW_FISH, it, Long::plus)
        }
    }

    return population
}

fun getNumberOfFishesAfterDummy(fishes: List<Int>, days: Int): Int {
    val mutFishes = fishes.toMutableList()
    for (`_` in 1..days) {
        val currentFishAmount = mutFishes.size - 1
        for (i in 0..currentFishAmount) {
            if (mutFishes[i] == 0) {
                mutFishes.add(8)
                mutFishes[i] = 6
            } else {
                mutFishes[i]--
            }
        }
    }
    return mutFishes.size
}


fun main() {
    val lanternFishes = File("src/input.txt")
        .readText()
        .trim()
        .split(",")
        .map { it.toInt() }

    println(getNumberOfFishesAfter(lanternFishes, 80))
    println(getNumberOfFishesAfter(lanternFishes, 256))
}