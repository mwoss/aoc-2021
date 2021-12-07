import java.io.File
import kotlin.math.*

enum class OptimizationMethod {
    CONSTANT,
    LINEAR,
}

fun findOptimalHorizontalPositionConstConsumption(objects: List<Int>): Int {
    // optimization for small lists has been omitted for simplicity and laziness :3
    return objects.sorted()[objects.size / 2]
}


fun findOptimalHorizontalPositionLinearConsumption(objects: List<Int>): Pair<Int, Int> {
    // optimization for small lists has been omitted for simplicity and laziness :3
    val avg = objects.average()
    return Pair(floor(avg).toInt(), ceil(avg).toInt())
}


fun getFuelSpentTowardsOptimalPosition(objects: List<Int>, method: OptimizationMethod): Int {
    return when (method) {
        OptimizationMethod.CONSTANT -> {
            val position = findOptimalHorizontalPositionConstConsumption(objects)
            return objects.fold(0) { acc, obj -> acc + abs(position - obj) }
        }
        OptimizationMethod.LINEAR -> {
            val (potentialPos1, potentialPos2) = findOptimalHorizontalPositionLinearConsumption(objects)
            val potentialSpent1 =
                objects.fold(0) { acc, obj -> acc + abs(potentialPos1 - obj) * (abs(potentialPos1 - obj) + 1) / 2 }
            val potentialSpent2 =
                objects.fold(0) { acc, obj -> acc + abs(potentialPos2 - obj) * (abs(potentialPos2 - obj) + 1) / 2 }
            if (potentialSpent1 < potentialSpent2) {
                potentialSpent1
            } else {
                potentialSpent2
            }
        }
    }
}


fun main() {
    val crabsPositions = File("src/input.txt").readText().trim().split(",").map { it.toInt() }

    println(getFuelSpentTowardsOptimalPosition(listOf(16, 1, 2, 0, 4, 2, 7, 1, 2, 14), OptimizationMethod.CONSTANT))
    println(getFuelSpentTowardsOptimalPosition(listOf(16, 1, 2, 0, 4, 2, 7, 1, 2, 14), OptimizationMethod.LINEAR))
    println(getFuelSpentTowardsOptimalPosition(crabsPositions, OptimizationMethod.CONSTANT))
    println(getFuelSpentTowardsOptimalPosition(crabsPositions, OptimizationMethod.LINEAR))

}