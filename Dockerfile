ARG ANDROID_HOME=/opt/google/android
ARG JAVA_HOME=/opt/oracle/openjdk


FROM mobiledevops/android-sdk-image:34.0.1 AS android

RUN rm -rf /opt/android-sdk-linux/emulator
RUN rm -rf /opt/android-sdk-linux/cmdline-tools
RUN rm -rf /opt/android-sdk-linux/extras
RUN rm -rf /opt/android-sdk-linux/platforms

# Disable Dependabot updates
FROM eclipse-temurin:17 AS java


FROM ccr.ccs.tencentyun.com/storezhang/ubuntu:24.04.24 AS builder
ARG ANDROID_HOME
ARG JAVA_HOME
COPY --from=android /opt/android-sdk-linux /docker/${ANDROID_HOME}
COPY --from=java /opt/java/openjdk /docker/${JAVA_HOME}
COPY flutter /docker/usr/local/bin/flutter


FROM ccr.ccs.tencentyun.com/storezhang/flutter:0.0.17

LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Flutter插件，此镜像只支持Android编译打包，如果需要其它平台的支持，请使用对应平台的镜像"

COPY --from=builder /docker /

RUN set -ex \
    \
    \
    \
    && apt update -y \
    && apt upgrade -y \
    \
    # 安装依赖库
    && apt install -y unzip \
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

ARG ANDROID_HOME
ARG JAVA_HOME

ENV ANDROID_HOME ${ANDROID_HOME}
ENV JAVA_HOME ${JAVA_HOME}
ENV PATH=${JAVA_HOME}/bin:${ANDROID_HOME}/bin:$PATH

# 配置依赖包缓存路径
ENV GRADLE_USER_HOME ${FLUTTER_CACHE}/gradle

# 默认为Android平台
ENV TYPE android
