VERSION=1.0.0
PREFIX=/usr/local
TARGET=gopom
MANPATH=${PREFIX}/man

$TARGET:
    go get
    go build -x -v -o ${TARGET}

clean:V:
    go clean -x

nuke:V:clean
    rm -f *.tgz *.sig

install:V:${TARGET}
    install -d ${TARGET} ${PREFIX}/bin
    install ${TARGET} ${PREFIX}/bin/${TARGET}
    install -d ${TARGET}.1 $MANPATH/man1
    install ${TARGET}.1 $MANPATH/man1/${TARGET}.1

uninstall:V:
    rm -f ${PREFIX}/bin/${TARGET}
    rm -f ${MANPATH}/man1/${TARGET}.1

lint:V:
    go vet

dist:V:nuke
    rm -f .*.swp
    cd .. && tar czf /tmp/${TARGET}-${VERSION}.tgz ${TARGET} && \
        cd ${TARGET} && mv /tmp/${TARGET}-${VERSION}.tgz .

