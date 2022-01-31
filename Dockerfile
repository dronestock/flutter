ARG FLUTTER_VERSION=2.8.1
ARG FLUTTER_HOME=/opt/google/flutter





FROM storezhang/alpine AS flutter

# 明确指定工作目录，防止后面运行命令出现文件或者目录找不到的问题
WORKDIR /opt


RUN apk update
RUN apk add axel

# 安装Flutter
ARG FLUTTER_VERSION
ARG FLUTTER_HOME


ENV FLUTTER_BIN_FILENAME flutter-${FLUTTER_VERSION}.tar.xz
ENV DOWNLOAD_URL https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_${FLUTTER_VERSION}-stable.tar.xz


RUN axel --num-connections 128 --output ${FLUTTER_BIN_FILENAME} --insecure ${DOWNLOAD_URL}
RUN tar -xzf ${JDK_BIN_FILENAME}
RUN mkdir -p ${FLUTTER_HOME}
RUN mv jdk-${FLUTTER_VERSION}/* ${FLUTTER_HOME}/





# 打包真正的镜像
FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成Flutter插件，支持测试、打包、发布等常规功能"


ARG FLUTTER_HOME


# 复制文件
COPY --from=flutter ${FLUTTER_HOME} ${FLUTTER_HOME}
COPY docker /
COPY flutter /bin


RUN set -ex \
    \
    \
    \
    # 安装依赖库
    && apk update \
    && apk --no-cache add libstdc++ gcompat gnupg \
    \
    # 解决找不到库的问题
    && LD_PATH=/etc/ld-musl-x86_64.path \
    && echo "/lib" >> ${LD_PATH} \
    && echo "/usr/lib" >> ${LD_PATH} \
    && echo "/usr/local/lib" >> ${LD_PATH} \
    && echo "${JAVA_HOME}/lib/default" >> ${LD_PATH} \
    && echo "${JAVA_HOME}/lib/j9vm" >> ${LD_PATH} \
    && echo "${JAVA_HOME}/lib/server" >> ${LD_PATH} \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/flutter \
    && chmod +x /usr/bin/gsk \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /bin/flutter


# 配置Flutter主目录
ENV FLUTTER_HOME ${FLUTTER_HOME}

# 将Flutter加入到系统路径中
ENV PATH=${FLUTTER_HOME}/bin:$PATH
