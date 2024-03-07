#include <iostream>




int main()
{



    std::string firstTeam;
    std::string firstTeamScore;

    std::string secondTeam;
    std::string secondTeamScore;

    std::string teamSolution;
    std::string scoreSolution;



    std::cin >> firstTeam >> firstTeamScore;
    std::cin >> secondTeam >> secondTeamScore;

    if (firstTeam.size() <= 12 && secondTeam.size() <= 12)
    {
         teamSolution = firstTeam  + " : "  + secondTeam;
        scoreSolution = firstTeamScore  + " : "  + secondTeamScore;


        while(teamSolution[13] != ':'){
            teamSolution = " " +  teamSolution;

            while(scoreSolution[13] != ':')
            {
                scoreSolution = " " + scoreSolution;
            }
        }
        std::cout << teamSolution << std::endl;
        std::cout << scoreSolution << std::endl;
    }
    return 0;

}
