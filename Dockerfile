ARG FLUTTER_HOME=/opt/google/flutter
ARG ANDROID_HOME=/opt/google/android
ARG JAVA_HOME=/opt/oracle/openjdk


FROM ghcr.io/cirruslabs/flutter:3.10.6 AS flutter

RUN rm -rf /sdks/flutter/dev
RUN rm -rf /sdks/flutter/examples

FROM mobiledevops/android-sdk-image:33.0.2 AS android

RUN rm -rf /opt/android-sdk-linux/emulator
RUN rm -rf /opt/android-sdk-linux/cmdline-tools
RUN rm -rf /opt/android-sdk-linux/extras
RUN rm -rf /opt/android-sdk-linux/platforms

# Disable Dependabot updates
FROM eclipse-temurin:11 AS java

FROM bitnami/git:2.41.0 AS git


FROM ccr.ccs.tencentyun.com/storezhang/ubuntu:23.04.17 AS builder
ARG FLUTTER_HOME
ARG ANDROID_HOME
ARG JAVA_HOME
COPY --from=flutter /sdks/flutter /docker/${FLUTTER_HOME}
COPY --from=android /opt/android-sdk-linux /docker/${ANDROID_HOME}
COPY --from=java /opt/java/openjdk /docker/${JAVA_HOME}
COPY --from=git /opt/bitnami/git/bin/git /docker/usr/local/bin/git
COPY flutter /docker/usr/local/bin/flutter


FROM ccr.ccs.tencentyun.com/storezhang/ubuntu:23.04.17

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
    # && apt upgrade -y \
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

ARG FLUTTER_HOME
ARG ANDROID_HOME
ARG JAVA_HOME

ENV FLUTTER_HOME ${FLUTTER_HOME}
ENV ANDROID_HOME ${ANDROID_HOME}
ENV JAVA_HOME ${JAVA_HOME}
ENV PATH=${JAVA_HOME}/bin:${FLUTTER_HOME}/bin:${ANDROID_HOME}/bin:$PATH

ENV FLUTTER_STORAGE_BASE_URL https://storage.flutter-io.cn
ENV PUB_HOSTED_URL https://pub.flutter-io.cn
ENV FLUTTER_GIT_URL https://gitee.com/mirrors/Flutter.git

# 配置依赖包缓存路径
ENV FLUTTER_CACHE /var/lib/flutter
ENV GRADLE_USER_HOME ${FLUTTER_CACHE}/gradle
ENV PUB_CACHE ${FLUTTER_CACHE}/pub

# 默认为Android平台
ENV TYPE android
