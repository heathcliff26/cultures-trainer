%global debug_package %{nil}

Name:           cultures-trainer
Version:        0
Release:        %autorelease
Summary:        Resource Cheat Trainer for Cultures

%global package_id io.github.heathcliff26.%{name}

License:        Apache-2.0
URL:            https://github.com/heathcliff26/cultures-trainer
Source:         %{url}/archive/refs/tags/v%{version}.tar.gz

BuildRequires: golang >= 1.24
BuildRequires: gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel libxkbcommon-devel wayland-devel

%global _description %{expand:
This Trainer is for the german version of the game. Hence the german resource names. Currently only tested with Northland, should also work with 8th Wonder of the World.}

%description %{_description}

%prep
%autosetup -n cultures-trainer-%{version} -p1

%build
make build

%install
install -D -m 755 bin/%{name} %{buildroot}/%{_bindir}/%{name}
install -D -m 644 packages/%{package_id}.desktop %{buildroot}/%{_datadir}/applications/%{package_id}.desktop
install -D -m 644 packages/%{package_id}.png %{buildroot}/%{_datadir}/icons/hicolor/256x256/apps/%{package_id}.png
install -D -m 644 %{package_id}.metainfo.xml %{buildroot}/%{_datadir}/metainfo/%{package_id}.metainfo.xml

%files
%license LICENSE
%doc README.md
%{_bindir}/%{name}
%{_datadir}/applications/%{package_id}.desktop
%{_datadir}/icons/hicolor/256x256/apps/%{package_id}.png
%{_datadir}/metainfo/%{package_id}.metainfo.xml

%changelog
%autochangelog
