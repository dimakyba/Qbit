using System;
using System.Linq;

class Program
{
    static void Main()
    {
        int k = int.Parse(Console.ReadLine());

        for (int t = 0; t < k; t++)
        {
            int n = int.Parse(Console.ReadLine());
            int[] array =  Console.ReadLine().Split(' ').Select(int.Parse).ToArray();

            bool hasSameSign = HasTwoAdjacentElementsWithSameSign(array);

            if (hasSameSign)
                Console.WriteLine("YES");
            else
                Console.WriteLine("NO");
        }
    }

    static bool HasTwoAdjacentElementsWithSameSign(int[] array)
    {
        for (int i = 0; i < array.Length - 1; i++)
        {
            if ((array[i] > 0 && array[i + 1] > 0) || (array[i] < 0 && array[i + 1] < 0))
            {
                return true;
            }
        }

        return false;
    }
}
