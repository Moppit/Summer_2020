import matplotlib.pyplot as plt

if __name__ == "__main__":
    # Read in each line of the data file
    data = {
        'date': [],
        'day': [],
        'pushups': [],
        'plank': [],
        'run': [],
        'comment':[]
    }
    f = open("challengeData.txt", "r")
    for line in f:
        if line != '':
            if line[0] != '!':
                toAdd = line.split(',')
                data['date'].append(toAdd[0])
                data['day'].append(int(toAdd[1]))
                data['pushups'].append(int(toAdd[2]))
                data['plank'].append(int(toAdd[3]))
                data['run'].append(float(toAdd[4]))
                data['comment'].append(toAdd[5])

    # Create individual graphs
    plt.suptitle('Pushups')
    plt.plot(data['day'], data['pushups'], 'r')
    plt.xlabel('Day')
    plt.ylabel('# pushups')
    plt.show()

    plt.suptitle('Planks')
    plt.plot(data['day'], data['plank'], 'b')
    plt.xlabel('Day')
    plt.ylabel('Plank Time (seconds)')
    plt.show()

    plt.suptitle('Running')
    plt.plot(data['day'], data['run'], 'g')
    plt.xlabel('Day')
    plt.ylabel('Distance (km)')
    plt.show()

    # Create combined graph
    plt.suptitle('All Metrics')
    line1,line2,line3 = plt.plot(data['day'], data['pushups'], 'r',
                                 data['day'], data['plank'], 'b',
                                 data['day'], data['run'], 'g')
    plt.xlabel('Day')
    plt.legend((line1, line2, line3), ('# pushups', 'plank (sec)', 'running (km)'))
    plt.show()
