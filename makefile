docker:
	docker build -t mfulki/user-service:latest ./user_service
	docker build -t mfulki/author-service:latest ./author_management
	docker build -t mfulki/book-service:latest ./book_management
	docker build -t mfulki/category-service:latest ./category_service