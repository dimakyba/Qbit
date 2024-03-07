using System;

class Program
{
    static void Main()
    {
        string nStr, mStr;
        double p;

        // Read p as a double
        string[] input = Console.ReadLine().Split();
        nStr = input[0];
        p = double.Parse(input[1]);
        mStr = input[2];

        p /= 100.0;

        // Convert n and m to decimal
        decimal n = decimal.Parse(nStr);
        decimal m = decimal.Parse(mStr);

        int i = 0;
        while (n < m)
        {
            n *= (decimal)(1 + p);
            i++;
        }

        Console.WriteLine(i);
    }
}
