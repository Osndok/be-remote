import QtQuick 2.4
import QtQuick.Layouts 1.1
import Ubuntu.Components 1.3

MainView {
    id: root
    objectName: 'mainView'
    applicationName: 'be-remote.reh'
    automaticOrientation: true

    width: units.gu(45)
    height: units.gu(75)

    Page {
        anchors.fill: parent

        header: PageHeader {
            id: header
            title: i18n.tr('BE Remote')
        }

        ColumnLayout {
            spacing: units.gu(2)
            anchors {
                margins: units.gu(2)
                top: header.bottom
                left: parent.left
                right: parent.right
                bottom: parent.bottom
            }

            Item {
                id: spacer
                Layout.fillHeight: true
            }

            Label {
                text: testvar.message
                horizontalAlignment: Label.AlignHCenter
                Layout.fillWidth: true

                /*
                    To export go functions they have to be in uppercase but in qml
                    the first letter of the function has to be lower case.
                    The qml<->go bridge takes care of that.
                */
            }

            Label {
                text: testvar.output
                horizontalAlignment: Label.AlignHCenter
                Layout.fillWidth: true
            }

            Button {
                text: 'Do it!'
                color: 'green'
                onClicked: testvar.getMessage()
                Layout.alignment: Qt.AlignCenter
            }

            Item {
                id: spacer2
                Layout.fillHeight: true
            }
        }
    }
}
