import QtQuick
import QtQuick.Layouts
import org.kde.plasma.components as PlasmaComponents
import org.kde.plasma.plasmoid

PlasmoidItem {
    Plasmoid.icon: "preferences-desktop"
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
            text: "Runtime status will be provided by org.xnix.Compatibility1."
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
        }
    }
}
