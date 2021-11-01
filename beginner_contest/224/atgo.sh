if [ -z "$1" ]; then
    echo "file name is required"
    exit 0
fi

if [ "1" = "$2" ]; then
    FILE=sample1.txt
    if [ ! -s $FILE ]; then
    echo "0 byte through: $FILE"
    else
    go run "$1.go" < $FILE
    fi
    exit 0
fi

if [ "2" = "$2" ]; then
    FILE=sample2.txt
    if [ ! -s $FILE ]; then
    echo "0 byte through: $FILE"
    else
    go run "$1.go" < $FILE
    fi
    exit 0
fi

if [ "3" = "$2" ]; then
    FILE=sample3.txt
    if [ ! -s $FILE ]; then
    echo "0 byte through: $FILE"
    else
    go run "$1.go" < $FILE
    fi
    exit 0
fi

if [ "4" = "$2" ]; then
    FILE=sample4.txt
    if [ ! -s $FILE ]; then
    echo "0 byte through: $FILE"
    else
    go run "$1.go" < $FILE
    fi
    exit 0
fi


FILE=sample1.txt
if [ ! -s $FILE ]; then
  echo "0 byte through: $FILE"
else
  go run "$1.go" < $FILE
fi

FILE=sample2.txt
if [ ! -s $FILE ]; then
  echo "0 byte through: $FILE"
else
  go run "$1.go" < $FILE
fi

FILE=sample3.txt
if [ ! -s $FILE ]; then
  echo "0 byte through: $FILE"
else
  go run "$1.go" < $FILE
fi

FILE=sample4.txt
if [ ! -s $FILE ]; then
  echo "0 byte through: $FILE"
else
  go run "$1.go" < $FILE
fi
