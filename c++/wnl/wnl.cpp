#include <stdio.h>
#include <conio.h>
#include <windows.h>
#include "time.h"

int key;
int mon[2][13] = {{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}, /*定义各个月的天数*/
                  {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}};
time_t lt;

int nowyear()
{
    struct tm *local;
    lt = time(NULL);
    local = localtime(&lt);
    return local->tm_year + 1900;
}
int nowmonth()
{
    struct tm *local;
    lt = time(NULL);
    local = localtime(&lt);
    return local->tm_mon + 1;
}
int nowday()
{
    struct tm *local;
    lt = time(NULL);
    local = localtime(&lt);
    return local->tm_mday;
}
int isleap(int year)
{
    int leap;
    leap = ((year % 4 == 0) && (year % 100 != 0) || (year % 400 == 0));
    return leap;
}
int isweek(int year, int month, int day)
{
    int i, j, week;
    int allday = 0;
    for (i = 0; i < year; i++)
    {
        if (1 == isleap(i))
        {
            allday = allday + 366;
        }
        else
        {
            allday = allday + 365;
        }
    }
    for (j = 1; j < month; j++)
    {
        allday = allday + mon[isleap(year)][j];
    }
    allday = allday + day - 1;
    week = (allday - 1) % 7;
    return week;
}
void checkdata(int year, int month)
{
    int i, z;
    int j = 0;
    printf(" %d年 %d月\n", year, month);
    printf(" ********************************************\n");
    printf("    日    一    二    三    四    五    六\n");
    printf(" ********************************************\n");
    z = isweek(year, month, 1);
    for (i = 0; i < z; i++)
    {
        printf("      ");
    }
    for (i = 0; i < mon[isleap(year)][month]; i++)
    {
        printf("%6d", i + 1);
        if ((0 == (z + i + 1) % 7) && ((i + 1) != mon[isleap(year)][month]))
        {
            j++;
            printf("\n");
        }
    }
    if (j == 3)
    {
        printf("\n\n\n\n");
    }
    if (j == 4)
    {
        printf("\n\n\n");
    }
    if (j == 5)
    {
        printf("\n\n");
    }
}
int strM()
{
    int year, month;
    printf("请输入要查询的年月\n>");
    scanf("%d%d", &year, &month);
    if (month > 12 || month < 1)
    {
        printf("非法日期!\n");
        printf("按任意键返回");
        getch();
        return 0;
    }
    printf("\n\n");
    key = 0;
    while (27 != key)
    {
        if (77 == key)
        {
            month++;
        }
        else if (75 == key)
        {
            month--;
        }
        if (month > 12)
        {
            month = month % 12;
            year = year + 1;
        }
        if (month < 1)
        {
            month = month + 12;
            year = year - 1;
        }
        system("cls");
        checkdata(year, month);
        printf(" 按左右方向键切换月份\n");
        printf(" 按ESC键退回");
        key = getch();
    }
    return 0;
}
int strY()
{
    int year, month = 1;
    
    printf("请输入要查询的年份\n>");
    scanf("%d", &year);
    printf("\n\n");
    key = 0;
    while (27 != key)
    {
        if (77 == key)
        {
            month++;
        }
        else if (75 == key)
        {
            month--;
        }
        else if (72 == key)
        {
            year++;
        }
        else if (80 == key)
        {
            year--;
        }
        if (month > 12)
        {
            month = month % 12;
            year = year + 1;
        }
        if (month < 1)
        {
            month = month + 12;
            year = year - 1;
        }
        system("cls");
        checkdata(year, month);
        printf(" 按方向键切换年月\n");
        printf(" 按ESC键退回");
        key = getch();
    }
    return 0;
}
int days()
{
    int y, m, d, date = 0;
    printf("请输入要查询的日期（年 月 日）\n>");
    scanf("%d%d%d", &y, &m, &d);
    if (m > 12 || m < 1 || d < 1 || d > mon[isleap(y)][m])
    {
        printf("非法日期!\n");
        printf("按任意键返回");
        getch();
        return 0;
    }
    if (y < nowyear())
    {
        for (int i = y; i < nowyear() - 1; i++)
        {
            if (1 == isleap(i))
            {
                date += 366;
            }
            else
            {
                date += 365;
            }
        }
        for (int i = m + 1; i <= 12; i++)
        {
            date += mon[isleap(y)][i];
        }
        for (int i = 0; i <= nowmonth() - 1; i++)
        {
            date += mon[isleap(y)][i];
        }
        date += mon[isleap(y)][m] - d;
        date += nowday();
    }
    else if (y > nowyear())
    {
        for (int i = nowyear() + 1; i < y; i++)
        {
            if (1 == isleap(i))
            {
                date += 366;
            }
            else
            {
                date += 365;
            }
        }
        for (int i = nowmonth() + 1; i <= 12; i++)
        {
            date += mon[isleap(y)][i];
        }
        for (int i = 0; i <= m - 1; i++)
        {
            date += mon[isleap(y)][i];
        }
        date += mon[isleap(y)][nowmonth()] - nowday();
        date += d;
    }
    else if (y == nowyear())
    {
        if (m > nowmonth())
        {
            for (int i = nowmonth() + 1; i < m; i++)
            {
                date += mon[isleap(y)][i];
            }
            date += mon[isleap(y)][nowmonth()] - nowday();
            date += d;
        }
        else if (m < nowmonth())
        {
            for (int i = m; i < nowmonth() - 1; i++)
            {
                date += mon[isleap(y)][i];
            }
            date += mon[isleap(y)][m] - d;
            date += nowday();
        }
        else if (m == nowmonth())
        {
            date += abs(d - nowday());
        }
    }
    printf("%d年%d月%d日距今%d天\n", y, m, d, date);
    printf("按任意键返回");
    getch();
    return 0;
}
int main()
{
    while (1)
    {
        int mkey;
        system("cls");
        printf("选择:\n1.年月->月历\n2.年份->年历\n3.日期->距今时间\n（按ESC退出）\n");
        mkey = getch();
        if ('1' == mkey)
        {
            system("cls");
            strM();
        }
        else if ('2' == mkey)
        {
            system("cls");
            strY();
        }
        else if ('3' == mkey)
        {
            system("cls");
            days();
        }
        else if (27 == mkey)
        {
            break;
        }
    }
    return 0;
}