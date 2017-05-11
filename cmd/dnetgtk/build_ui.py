#!/bin/env python3


import subprocess
import os.path
import glob

WIDGETS = [
    './common_windgets/key_cert_editor',
    './common_windgets/text_viewer',
    './common_windgets/progress_window'
    ]

MODULES = [
    './applications/builtin_ownkeypair',
    './applications/builtin_owntlscert'
]

DNETGETK = [
    "."
]

ALL_DIRS = sorted(WIDGETS + MODULES + DNETGETK)

print("generating ui code")


def rm_backup_files(path):
    res = glob.glob(os.path.join(path, '*~'))
    for i in res:
        print("        rm {}".format(os.path.basename(i)))
        os.unlink(i)


def makedirsui(path):
    o_path = path
    path = os.path.abspath(path)
    base = os.path.basename(path)
    dirn = os.path.dirname(path)
    if o_path != '.':
        pkg_name = base
    else:
        pkg_name = 'main'

    print("    {}".format(o_path))

    rm_backup_files(os.path.join(o_path, 'ui'))

    cmd = ["go-bindata", "-o", "ui_bindata.go", "-pkg", pkg_name, "ui"]

    print("        {}".format(' '.join(cmd)))

    p = subprocess.Popen(cmd, cwd=path)
    res = p.wait()

    return res

for i in ALL_DIRS:
    res = makedirsui(i)
    if res != 0:
        print("   error")
        exit(1)

print("success")
exit(0)
