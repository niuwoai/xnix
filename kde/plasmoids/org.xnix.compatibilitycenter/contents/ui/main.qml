import QtQuick
import QtQuick.Layouts
import org.kde.plasma.components as PlasmaComponents
import org.kde.plasma.plasmoid

PlasmoidItem {
    Plasmoid.icon: "preferences-desktop"
    property string runtimeModelCommand: "xnix-kde-center-model"
    readonly property string centerPreviewCommand: "kde-center-page-preview"
    readonly property string guiEvidenceCountField: "known_app_gui_evidence_count"
    readonly property string guiEvidenceCardsField: "known_app_gui_evidence_cards"
    readonly property string guiEvidenceSource: "wine-guest-gui-smoke"
    readonly property string guiEvidenceKind: "known-application-gui-smoke"
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
            text: "Real Windows GUI evidence"
            font.bold: true
            Layout.fillWidth: true
        }
        PlasmaComponents.Label {
            text: "This page expects kde-center-page-preview to provide known_app_gui_evidence_count and known_app_gui_evidence_cards from wine-guest-gui-smoke."
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
        }
        PlasmaComponents.Label {
            text: "A passing GUI card means a local Windows .exe was copied into the QEMU guest, launched by Wine, and observed as an X11 window."
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
