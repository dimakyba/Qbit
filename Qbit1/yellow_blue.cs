using System;
using System.Text;

namespace Program
{
    class Program
    {

        static void Main(string[] args)
        {
            Console.OutputEncoding = Encoding.UTF8;
            string team1 = Console.ReadLine();
            string s1 = Console.ReadLine();
            string team2 = Console.ReadLine();
            string s2 = Console.ReadLine();
            string line1 = ($"{team1} : {team2}");
            string line2 = ($"{s1} : {s2}");
            while (line1.IndexOf(":") < 13)
            {
                line1 = line1.Insert(0, " ");
            }
            while (line2.IndexOf(":") < 13)
            {
                line2 = line2.Insert(0, " ");

            }

            Console.WriteLine(line1);
            Console.WriteLine(line2);
        }

    }
}
