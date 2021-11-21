FROM ubuntu:latest
RUN useradd -u 10001 scratchuser

FROM scratch
COPY ./battlesnake_arm64 /battlesnake
COPY --from=0 /etc/passwd /etc/passwd
EXPOSE 8080
CMD ["/battlesnake"]
