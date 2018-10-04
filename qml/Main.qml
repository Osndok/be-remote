import QtQuick 2.4
import QtQuick.Layouts 1.1
import Ubuntu.Components 1.3

MainView {
    id: root
    objectName: 'mainView'
    applicationName: 'be_remote.robert'
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

            Label
            {
                text: "Message"
            }

            TextArea
            {
                id: myTextArea
                focus: true
                text: testvar.message
                horizontalAlignment: Label.AlignHCenter
                Layout.fillWidth: true
                Component.onCompleted: myTextArea.forceActiveFocus()

                /*
                    To export go functions they have to be in uppercase but in qml
                    the first letter of the function has to be lower case.
                    The qml<->go bridge takes care of that.
                */
            }

            RowLayout
            {
            Layout.alignment: Qt.AlignCenter

            Button
            {
                text: 'Cancel'
                onClicked: testvar.doCancel()
                //Layout.alignment: Qt.AlignCenter
            }

            Button
            {
                text: 'Note'
                //color: 'green'
                onClicked: testvar.doNote()
            }

            Button
            {
                text: 'Todo'
                color: 'green'
                onClicked: testvar.doTodo()
                //Layout.alignment: Qt.AlignCenter
            }
            }

            Label
            {
                text: testvar.output
                horizontalAlignment: Label.AlignHCenter
                Layout.fillWidth: true
            }

            Item {
                id: spacer2
                Layout.fillHeight: true
            }
        }
    }
}

