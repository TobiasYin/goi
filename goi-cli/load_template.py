import os
import sys


def load():
    if not os.path.exists("template"):
        print("error, template not found!")
        sys.exit(-1)
    return pack_dir("./", "template")


def pack_file(base, name):
    temp = """File{{Content: []byte{{{}}},Name: "{}",IsFile: true,}},"""
    with open(os.path.join(base, name), "rb") as f:
        content = f.read()
    res = [str(i) for i in content]
    # content = content.replace("\\n", "\\\\n")
    # content = content.replace("\n", "\\n")
    # content = content.replace("\"", "\\\"")
    # content = content.replace("\\0", "\\\\0")
    res = temp.format(",".join(res), name)
    return res


def pack_dir(base, name):
    temp = """File{{IsDir: true,Name: "{}",Children: []File{{{}}},}},"""
    path = os.path.join(base, name)
    l = os.listdir(path)
    children = []
    for i in l:
        p = os.path.join(path, i)
        if os.path.isdir(p):
            children.append(pack_dir(path, i))
        elif os.path.isfile(p):
            children.append(pack_file(path, i))

    res = temp.format(name, " ".join(children))
    return res


def main():
    file_temp = """package inline\nfunc init()  {{Root = {}}}"""
    content = file_temp.format(load()[:-1])
    with open(os.path.join("inline", "init.go"), "w") as f:
        f.write(content)


if __name__ == "__main__":
    main()
