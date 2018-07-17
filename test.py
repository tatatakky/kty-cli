import sys

def option_judge(option1, option2, arguments):
    if option1 in arguments or option2 in arguments:
        info = ''
        idx = arguments.index(option1)+1

        if option1 == '-s' or option2 == '-sns':
            return arguments[idx]
        else:
            for i in range(idx,len(arguments)):
                if arguments[i].startswith('-'):
                    break
                info += arguments[i] + ' '
            return info[:-1]

def parser():
    info = []
    arguments = sys.argv
    arguments.pop(0)
#    info.append(option_judge('-m', '-message', arguments))
    info.append(option_judge('-u', '-user', arguments))
    info.append(option_judge('-m', '-message', arguments))
    info.append(option_judge('-s', '-sns', arguments))

    return info

if __name__ == '__main__':

    infomation = parser()

    for i in range(len(infomation)):
        if i == 0:
            print("To      : {}".format(infomation[i]))
        elif i == 1:
            print("Message : {}".format(infomation[i]))
        else:
            print("sns     : {}".format(infomation[i]))


