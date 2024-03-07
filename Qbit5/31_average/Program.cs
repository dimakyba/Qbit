using System;
using System.Linq;

class Program
{
    static void Main()
    {
        int n = int.Parse(Console.ReadLine());
        int[] array = Console.ReadLine().Split(' ').Select(int.Parse).ToArray();

        double result = AverageOddElements(array);
        Console.WriteLine(result.ToString("0.000"));
    }

    static double AverageOddElements(int[] arr)
    {
        double sum = 0;
        int count = 0;

        for (int i = 1; i < arr.Length; i += 2)
        {
            if (arr[i] % 2 != 0)
            {
                sum += arr[i];
                count++;
            }
        }

        if (count == 0)
        {
            return 0;
        }

        return sum / count;
    }
}
