using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;

namespace NorthwindLib
{
    public class Northwind
        : DbContext
    {
        public DbSet<Category> Categories { get; set; }
        public DbSet<Product> Products { get; set; }

        public Northwind(DbContextOptions<Northwind> options)
        : base(options)
        {

        }
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            base.OnModelCreating(modelBuilder);
            modelBuilder.Entity<Category>()
            .Property(c => c.CategoryName)
            .IsRequired()
            .HasMaxLength(20);

            modelBuilder.Entity<Category>()
            .HasMany(c => c.Products)
            .WithOne(p => p.Category);

            modelBuilder.Entity<Product>()
            .Property(c => c.ProductName)
            .IsRequired()
            .HasMaxLength(40);

            modelBuilder.Entity<Product>()
            .HasOne(p => p.Category)
            .WithMany(c => c.Products);
        }
    }
}