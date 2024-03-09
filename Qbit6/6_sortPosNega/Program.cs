using System;
using System.Linq;

class Program
{
  static void Main()
  {
    int n = int.Parse(Console.ReadLine());

    int[] arr = Console.ReadLine().Split().Select(int.Parse).ToArray();

    var negatives = arr.Where(x => x < 0).OrderBy(x => x);
    var zeros = arr.Where(x => x == 0).OrderBy(x => x);
    var positives = arr.Where(x => x > 0).OrderByDescending(x => x);

    Console.WriteLine(string.Join(" ", negatives.Concat(zeros).Concat(positives)));
    Console.ReadLine();
  }
}
