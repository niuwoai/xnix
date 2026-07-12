import QtQuick
import QtQuick.Layouts
import org.kde.plasma.components as PlasmaComponents
import org.kde.plasma.plasmoid

PlasmoidItem {
    Plasmoid.icon: "preferences-desktop"
    property string runtimeModelCommand: "xnix-kde-center-model"
    compactRepresentation: PlasmaComponents.ToolButton {
        icon.name: "preferences-desktop"
        text: "Compatibility"
        onClicked: Plasmoid.expanded = !Plasmoid.expanded
    }
    fullRepresentation: ColumnLayout {
        implicitWidth: 300
        spacing: 8

        PlasmaComponents.Label {
            text: "Xnix Compatibility Center"
            font.bold: true
        }
        PlasmaComponents.Label {
            text: "Runtime-backed summary is provided by xnix-kde-center-model."
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
        }
        PlasmaComponents.Label {
            text: "Desktop integration stays read-only here; install, launch, repair, snapshot, and restore actions remain Runtime requests."
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
        }
    }
}
