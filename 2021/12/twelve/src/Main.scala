import scala.collection.mutable
import scala.collection.mutable.ArrayBuffer




//noinspection DuplicatedCode
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
    val completedPaths = ArrayBuffer[ArrayBuffer[String]]()

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

  def part2(): Unit = {
    val caves = readCaves()
    val completedPaths = ArrayBuffer[ArrayBuffer[String]]()
    def canMove(visited: mutable.HashMap[String, Int], name: String, hasUtilizedException: Boolean): Int = {
      val smallCave = Character.isLowerCase(name(0))
      if (!smallCave) {
        return 1
      }

      if (visited.getOrElse(name, 0) == 0) {
        return 1
      }

      if (!hasUtilizedException && (name != "start" && name != "end")) {
        return 2
      }

      0
    }


    def traverse(currentCaveName: String, visited: mutable.HashMap[String, Int], currentPath: ArrayBuffer[String], hasUtilizedException: Boolean): Unit = {
      if (currentCaveName == "end") {
        currentPath += currentCaveName
        completedPaths += currentPath
        return
      }
      visited(currentCaveName) = visited.getOrElse(currentCaveName, 0) + 1
      currentPath += currentCaveName
      val nextCaves = caves.getOrElse(currentCaveName, ArrayBuffer[String]())
      for (cave <- nextCaves) {
        val moveStatus = canMove(visited, cave, hasUtilizedException)
        if (moveStatus >= 1) {
          if (moveStatus == 2) traverse(cave, visited, currentPath.clone(), true)
          else traverse(cave, visited, currentPath.clone(), hasUtilizedException)
        }
      }
      visited(currentCaveName) = visited.getOrElse(currentCaveName, 0) - 1
      currentPath.remove(currentPath.length - 1)
    }

    traverse("start", mutable.HashMap[String, Int](), ArrayBuffer[String](), false)
    println(completedPaths.length)
  }

part2()

}
