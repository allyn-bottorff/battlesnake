FROM scratch
COPY ./battlesnake_arm64 /battlesnake
EXPOSE 8080
CMD ["/battlesnake"]
