import scala.util.Using
import scala.collection.mutable


object Main {
  def split(l: Array[String], i: String): Array[Array[String]] = {
    l match {
      case Array() => Array()
      case _ =>
        val (h, t) = l.span(a => a != i)
        Array(h) ++ split(t.drop(1), i)
    }
  }

  def parseCoordinatePairLine(line: String): (Int, Int) = {
    val splitLine = line.split(',')

    (splitLine(0).toInt, splitLine(1).toInt)
  }

  def parseFoldInstructionLine(line: String): (Char, Int) = {
    val splitLine = line.split('=')


    (splitLine(0).last, splitLine(1).toInt)
  }

  def readPaper(): Set[(Int, Int)] = {
    val lines = Using(scala.io.Source.fromFile("input.txt")) {
      _.mkString.split('\n')
    }.get
    val coordinatePairLines = split(lines, "")(0)
    val coordinatePairs = for (x <- coordinatePairLines) yield parseCoordinatePairLine(x)
    coordinatePairs.toSet
  }

  def readFoldInstructions(): Array[(Char, Int)] = {
    val lines = Using(scala.io.Source.fromFile("input.txt")) {
      _.mkString.split('\n')
    }.get

    val foldInstructionsLines = split(lines, "")(1)
    val foldInstructions = for (x <- foldInstructionsLines) yield parseFoldInstructionLine(x)
    foldInstructions
  }

  def processXFold(axis: Int, dots: mutable.Set[(Int, Int)]): mutable.Set[(Int, Int)] = {
    val foldedPoints = dots.filter(_._1 > axis)
    for (dot <- foldedPoints) {
      dots.remove(dot)
      val newXPosition = 2 * axis - dot._1
      dots.add((newXPosition, dot._2))
    }
    dots
  }

  def processYFold(axis: Int, dots: mutable.Set[(Int, Int)]): mutable.Set[(Int, Int)] = {
    val foldedPoints = dots.filter(_._2 > axis)
    for (dot <- foldedPoints) {
      dots.remove(dot)
      val newYPosition = 2 * axis - dot._2
      dots.add((dot._1, newYPosition))
    }
    dots
  }

  def processFold(instruction: (Char, Int), dots: mutable.Set[(Int, Int)]):  mutable.Set[(Int, Int)] = {
    if (instruction._1 == 'x') {
      val result = processXFold(instruction._2, dots)
      return result
    }
     else if (instruction._1 == 'y') {
      val result = processYFold(instruction._2, dots)
      return result
    }

    mutable.Set[(Int, Int)]()
  }

  def part1(): Unit = {
    val dots = readPaper()
    val foldInstructions = readFoldInstructions()
    val reducedDots = processFold(foldInstructions(0), dots.to(mutable.Set))
    println(reducedDots.knownSize)
  }

  def part2(): Unit = {
    var dots = readPaper().to(mutable.Set)
    val foldInstructions = readFoldInstructions()
    for (fold <- foldInstructions) {
      println(dots)
      dots = processFold(fold, dots)
    }
    for (i <- dots.map(_._2).min to dots.map(_._2).max) {
      for (j <- dots.map(_._1).min to dots.map(_._1).max) {
        if (dots.contains((j, i))) {
          print("X")
        } else {
          print(" ")
        }
      }
      print("\n")
    }
  }

  def main(args: Array[String]): Unit = {
    part2()
  }
}