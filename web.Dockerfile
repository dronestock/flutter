FROM ccr.ccs.tencentyun.com/storezhang/flutter:0.0.67

LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Flutter插件，此镜像只支持Web平台编译打包，如果需要其它平台的支持，请使用对应平台的镜像"

COPY flutter /usr/local/bin/flutter

RUN set -ex \
    \
    \
    \
    && apt update -y \
    && apt upgrade -y \
    \
    \
    \
    # 增加执行权限
    && chmod +x /usr/local/bin/flutter \
    \
    \
    \
    # 清理镜像，减少无用包
    && apt autoremove -y \
    && rm -rf /var/lib/apt/lists/* \
    && apt autoclean

# 执行命令
ENTRYPOINT /usr/local/bin/flutter

# 默认为Web平台
ENV TYPE web
