import java.io.File

const val MAX_HEIGHT = 9

fun calculateRiskLevel(heights: Array<Array<Int>>): Int {
    var riskLevel = 0
    for (i in heights.indices) {
        for (j in heights[0].indices) {
            val up = if (i - 1 >= 0) heights[i - 1][j] else MAX_HEIGHT
            val down = if (i + 1 < heights.size) heights[i + 1][j] else MAX_HEIGHT
            val right = if (j + 1 < heights[0].size) heights[i][j + 1] else MAX_HEIGHT
            val left = if (j - 1 >= 0) heights[i][j - 1] else MAX_HEIGHT
            val current = heights[i][j]

            if (current < up && current < down && current < right && current < left) {
                riskLevel += (1 + current)
            }
        }
    }
    return riskLevel
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
}