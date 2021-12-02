package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bagasunix/go_blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// Create blog
	fmt.Println("Creating the blog")
	blog := blogpb.Blog{
		AuthorId: "Aldino Pratama",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: &blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been creater: %v", createBlogRes)
}
