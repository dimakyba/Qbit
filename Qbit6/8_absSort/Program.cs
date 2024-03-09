using System;
using System.Linq;

class Program
{
    static void Main()
    {
        int N = int.Parse(Console.ReadLine());
        int[] scores = Console.ReadLine().Split().Select(int.Parse).ToArray();

        var sortedScores = scores.Select((score, index) => new { Index = index + 1, Score = score })
                                .OrderByDescending(x => x.Score)
                                .ToArray();

        foreach (var participant in sortedScores)
        {
            Console.Write(participant.Index + " ");

        }
    }
}
