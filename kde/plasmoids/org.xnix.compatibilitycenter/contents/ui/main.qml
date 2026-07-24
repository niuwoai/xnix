import QtQuick
import QtQuick.Layouts
import org.kde.plasma.components as PlasmaComponents
import org.kde.plasma.plasmoid

PlasmoidItem {
    Plasmoid.icon: "preferences-desktop"
    property string runtimeModelCommand: "xnix-kde-center-model"
    readonly property string centerPreviewCommand: "kde-center-page-preview"
    readonly property string guiEvidenceCountField: "known_app_gui_evidence_count"
    readonly property string ownerControlledGuiEvidenceCountField: "known_app_owner_controlled_gui_evidence_count"
    readonly property string ownerManagedCopyVerifiedCountField: "known_app_owner_managed_copy_verified_count"
    readonly property string guiEvidenceCardsField: "known_app_gui_evidence_cards"
    readonly property string guiEvidenceSource: "wine-guest-gui-smoke"
    readonly property string guiEvidenceKind: "known-application-gui-smoke"
    readonly property var guiEvidenceCardFields: [
        "display_name",
        "smoke_status",
        "compatibility_state",
        "center_card_state",
        "execution_evidence_recorded",
        "staged_launcher_verified",
        "owner_controlled_runtime_launch_verified",
        "owner_managed_copy_verified",
        "owner_service_call_ready",
        "owner_evidence_handoff_ready",
        "owner_evidence_relative_path",
        "runtime_dispatch_verified",
        "primary_action_id",
        "primary_action_kind",
        "primary_action_label",
        "desktop_callable_route",
        "desktop_callable_runtime_method",
        "desktop_callable_execution_type",
        "desktop_dbus_method",
        "desktop_evidence_handle_forwarded",
        "kde_forwarded_argument_kind",
        "kde_forwarded_arguments",
        "owner_service_args_exposed_to_kde"
    ]
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
            text: "This page expects kde-center-page-preview to provide known_app_gui_evidence_count, known_app_owner_controlled_gui_evidence_count, known_app_owner_managed_copy_verified_count, and known_app_gui_evidence_cards from wine-guest-gui-smoke."
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
        }
        PlasmaComponents.Label {
            text: "The KDE read model renders display_name, smoke_status, compatibility_state, center_card_state, execution_evidence_recorded, owner_controlled_runtime_launch_verified, owner_managed_copy_verified, owner_evidence_handoff_ready, runtime_dispatch_verified, primary_action_id, primary_action_kind, primary_action_label, desktop_callable_route, desktop_dbus_method, and desktop_evidence_handle_forwarded from each safe card."
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
        }
        PlasmaComponents.Label {
            text: "When a card is handoff-ready, KDE forwards only the evidence-relative-path handle to the Runtime D-Bus action; owner service arguments and backend paths stay hidden."
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
        }
        PlasmaComponents.Label {
            text: "A passing owner-controlled GUI card means the Runtime owner service invoked the managed launcher, copied the Windows .exe into the QEMU guest when needed, launched it by Wine, and observed it as an X11 window."
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
