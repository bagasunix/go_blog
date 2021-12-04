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
	blogId := createBlogRes.GetBlog().GetId()

	// Reading the blog
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "61ab202ea8ba0a46d1c049b0"})
	if err2 != nil {
		fmt.Sprintf("Error happened while reading: %v", err2)
	}
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v", readBlogErr)
	}

	fmt.Printf("Blog was read: %v", readBlogRes)
}
