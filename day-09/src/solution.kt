import java.io.File
import java.util.*

private const val MAX_HEIGHT = 9

class Point(
    val current: Int,
    private val up: Int,
    private val down: Int,
    private val right: Int,
    private val left: Int
) {
    companion object {
        fun of(heights: Array<Array<Int>>, x: Int, y: Int): Point {
            val up = if (x - 1 >= 0) heights[x - 1][y] else MAX_HEIGHT
            val down = if (x + 1 < heights.size) heights[x + 1][y] else MAX_HEIGHT
            val right = if (y + 1 < heights[0].size) heights[x][y + 1] else MAX_HEIGHT
            val left = if (y - 1 >= 0) heights[x][y - 1] else MAX_HEIGHT
            return Point(heights[x][y], up, down, right, left)
        }
    }

    fun isLowPoint(): Boolean {
        if (current < up && current < down && current < right && current < left) {
            return true
        }
        return false
    }
}


fun calculateRiskLevel(heights: Array<Array<Int>>): Int {
    var riskLevel = 0
    for (x in heights.indices) {
        for (y in heights[0].indices) {
            val point = Point.of(heights, x, y)
            if (point.isLowPoint()) {
                riskLevel += (1 + point.current)
            }
        }
    }
    return riskLevel
}

fun calculateMultiplicationOfLargestBasins(heights: Array<Array<Int>>, n: Int): Int {
    val basins = mutableListOf<Int>()

    for (x in heights.indices) {
        for (y in heights[0].indices) {
            val point = Point.of(heights, x, y)
            if (point.isLowPoint()) {
                basins.add(getBasinSizeAt(heights, x, y))
            }
        }
    }

    return basins.sorted().takeLast(n).reduce { acc, basin -> acc * basin }
}

private fun getBasinSizeAt(heights: Array<Array<Int>>, x: Int, y: Int): Int {
    var basinSize = 0
    val visited = mutableSetOf<Pair<Int, Int>>()
    val queue = LinkedList<Pair<Int, Int>>()

    queue.add(Pair(x, y))
    visited.add(Pair(x, y))

    while (queue.isNotEmpty()) {
        val (xp, yp) = queue.pop()

        basinSize++

        val up = Pair(xp - 1, yp)
        val down = Pair(xp + 1, yp)
        val left = Pair(xp, yp - 1)
        val right = Pair(xp, yp + 1)

        if (xp - 1 >= 0 && heights[xp - 1][yp] < MAX_HEIGHT && up !in visited) {
            queue.add(up)
            visited.add(up)
        }
        if (xp + 1 < heights.size && heights[xp + 1][yp] < MAX_HEIGHT && down !in visited) {
            queue.add(Pair(xp + 1, yp))
            visited.add(down)
        }
        if (yp - 1 >= 0 && heights[xp][yp - 1] < MAX_HEIGHT && left !in visited) {
            queue.add(Pair(xp, yp - 1))
            visited.add(left)
        }
        if (yp + 1 < heights[0].size && heights[xp][yp + 1] < MAX_HEIGHT && right !in visited) {
            queue.add(Pair(xp, yp + 1))
            visited.add(right)
        }
    }

    return basinSize
}

fun main() {
    val heights = File("src/input.txt")
        .readLines()
        .map { lines ->
            lines.split("")
                .mapNotNull { it.toIntOrNull() }
                .toTypedArray()
        }
        .toTypedArray()

    println(calculateRiskLevel(heights))
    println(calculateMultiplicationOfLargestBasins(heights, 3))
}