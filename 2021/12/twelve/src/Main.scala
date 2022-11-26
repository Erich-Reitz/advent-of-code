import scala.collection.mutable
import scala.collection.mutable.ArrayBuffer




object Main extends App {
  def readCaves(): mutable.HashMap[String, ArrayBuffer[String]] = {
    val caves: mutable.HashMap[String, ArrayBuffer[String]] = mutable.HashMap()
    val lines = scala.io.Source.fromFile("input.txt").mkString.split('\n')
    for (connection: String <- lines) {
      val connectionParts = connection.split('-')
      val origin: String = connectionParts(0)
      val next = connectionParts(1)
      caves(origin) = caves.getOrElse(origin, ArrayBuffer[String]()) += next
      caves(next) = caves.getOrElse(next, ArrayBuffer[String]()) += origin
    }
    caves
  }

  def part1(): Unit = {
    val caves = readCaves()
    var completedPaths = ArrayBuffer[ArrayBuffer[String]]()

    def traverse(currentCaveName: String, visited: mutable.Set[String], currentPath: ArrayBuffer[String]): Unit = {
      if (currentCaveName == "end") {
        currentPath += currentCaveName
        completedPaths += currentPath
        return
      }
      visited.add(currentCaveName)
      currentPath += currentCaveName
      val nextCaves = caves.getOrElse(currentCaveName, ArrayBuffer[String]())
      for (cave <- nextCaves) {
        val smallCave = Character.isLowerCase(cave(0))
        if (!smallCave || (smallCave && !visited.contains(cave))) {
          traverse(cave, visited, currentPath.clone())
        }
      }
      visited.remove(currentCaveName)
      currentPath.remove(currentPath.length - 1)
    }

    traverse("start", mutable.Set[String](), ArrayBuffer[String]())
    println(completedPaths.length)
  }

  part1()

}
