# Day 01 — Go Boot camp

## Comparing Incomparable

💡 [Tap here](https://new.oprosso.net/p/4cb31ec3f47a4596bc758ea1861fb624) **to leave your feedback on the project**. It's anonymous and will help our team make your educational experience better. We recommend completing the survey immediately after the project.

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00: Reading](#exercise-00-reading)
5. [Chapter V](#chapter-v) \
    5.1. [Exercise 01: Assessing Damage](#exercise-01-assessing-damage)
6. [Chapter VI](#chapter-vi) \
    6.1. [Exercise 02: Afterparty](#exercise-02-afterparty)


<h2 id="chapter-i" >Chapter I</h2>
<h2 id="general-rules" >General rules</h2>

- Your programs should not exit unexpectedly (give an error on valid input). If this happens, your project will be considered non-functional and will receive a 0 in the evaluation.
- We encourage you to create test programs for your project, even though this work doesn't have to be submitted and won't be judged. This will allow you to easily test your work and the work of your peers. You will find these tests particularly useful during your defence. In fact, you are free to use your own tests and/or those of the peer you are evaluating during your defence.
- Submit your work to your allocated git repository. Only work in the git repository will be marked.
- If your code uses external dependencies, it should use [Go Modules](https://go.dev/blog/using-go-modules) to manage them.

<h2 id="chapter-ii" >Chapter II</h2>
<h2 id="rules-of-the-day" >Rules of the day</h2>

- You should only submit `*.go` files and (in the case of external dependencies) `go.mod` + `go.sum`.
- Your code for this task should be buildable with just `go build`.

<h2 id="chapter-iii" >Chapter III</h2>
<h2 id="intro" >Intro</h2>

There are many popular data formats in the world of programming, and Go in particular. But it's very likely that you'll come across one of them — XML or JSON. Many, many APIs use JSON and/or XML to encode structured data.

And... sometimes even bakeries use it to store recipes. For example, the old famous bakery in Villariba used to use good old XML to store the list of cake recipes. If we take a look at that piece of the database, it looks something like this:

```xml
<recipes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>40 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>3</itemcount>
                <itemunit>cups</itemunit>
            </item>
            <item>
                <itemname>Vanilla extract</itemname>
                <itemcount>1.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>7</itemcount>
                <itemunit></itemunit> <!-- itemunit may be empty  -->
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
                <itemunit>pieces</itemunit>
            </item>
            <!-- Here can be more ingredients  -->
        </ingredients>
    </cake>
    <cake>
        <name>Blueberry Muffin Cake</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Baking powder</itemname>
                <itemcount>3</itemcount>
                <itemunit>teaspoons</itemunit>
            </item>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>0.5</itemcount>
                <itemunit>cup</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>cup</itemunit>
            </item>
            <!-- Here can be more ingredients  -->
        </ingredients>
    </cake>
    <!-- Here can be more cakes  -->
</recipes>
```

Life was great and simple until the owner of the bakery found out that in the next village, Villabajo, there lived a dirty impostor who had managed to steal his recipes! To get away with his trick, he even used a different data format to store them and hide from justice!

```json
{
  "cake": [
    {
      "name": "Red Velvet Strawberry Cake",
      "time": "45 min",
      "ingredients": [
        {
          "ingredient_name": "Flour",
          "ingredient_count": "2",
          "ingredient_unit": "mugs"
        },
        {
          "ingredient_name": "Strawberries",
          "ingredient_count": "8"  // ingredient_unit is not even here!
        },
        {
          "ingredient_name": "Coffee Beans",
          "ingredient_count": "2.5",
          "ingredient_unit": "tablespoons"
        },
        {
          "ingredient_name": "Cinnamon",
          "ingredient_count": "1"
        }
      ]
    },
    {
      "name": "Moonshine Muffin",
      "time": "30 min",
      "ingredients": [
        {
          "ingredient_name": "Brown sugar",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        },
        {
          "ingredient_name": "Blueberries",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        }
      ]
    }
  ]
}
```

He couldn't help but notice that the thief had not only stolen his recipes, but had also altered some of them. Some ingredients were missing, numbers were changed, units were renamed. So he prepared for revenge!


<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Exercise 00: Reading</h3>

First of all, he had to learn how to read the database. The owner already had a CLI, so he decided that reading the file should be straightforward, so both these should work (files can be distinguished by an extension, for simplicity):

`~$ ./readDB -f original_database.xml`
`~$ ./readDB -f stolen_database.json`

Not only that, but he also decided that it shouldn't be too difficult to read both files using the same interface, which he called `DBReader`. That is, reading different formats means that we have different *implementations* of the same `DBReader` interface, which should spit out the same object types as a result, whether it's reading from the original database or the stolen one. Right, his idea is to choose the appropriate implementation based on a file extension.

So you're going to have to help him with that. Think about what kinds of objects are in these databases and how they can be represented in code. Then write an interface `DBReader` and two implementations of it — one for reading JSON and one for reading XML. Both should return an object of the same type as a result.

To check that his idea works, have the code print the JSON version of the database when it's reading from XML, and vice versa. Both XML and JSON fields should be indented by 4 spaces ("pretty-printing").

<h2 id="chapter-v" >Chapter V</h2>
<h3 id="ex01">Exercise 01: Assessing Damage</h3>

OK, so now the owner has decided to compare the databases. You've seen that the stolen database has modified versions of the same recipes, so there are several possible cases:

1) New cake added or old cake removed.
2) Different cooking time for the same cake.
3) A new ingredient is added or removed for the same cake. *Important:* The order of the ingredients doesn't matter. Only the names do.
4) The number of units for the same ingredient has changed.
5) The unit used to measure the ingredient has changed.
6) An ingredient unit is missing or has been added.

On a quick scan of the database, the owner also noticed that nobody bothered to rename the cakes or ingredients (maybe just added some new ones), so you can assume that the names are the same in both databases.

Your application should work like this:

`~$ ./compareDB --old original_database.xml --new stolen_database.json`

It should work with both formats (JSON and XML) for the original AND the new database, reusing the code from Exercise 00.

The output should look like this (same cases as above):

```
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
```

<h2 id="chapter-vi" >Chapter VI</h2>
<h3 id="ex02">Exercise 02: Afterparty</h3>

As the owner of the Villariba bakery searched through the database, he suddenly realised — this guy is brilliant! Some of the recipes had been greatly improved from the old version, and these new ideas were truly creative! He rushed to Villabajo and found the man who, as he first thought, had stolen his most precious inheritance.

...That same evening in the tavern, two old bakers were hugging, drinking and laughing so hard that it could be heard in both villages. They had become best friends over the past few hours, and each of them was delighted to have found the person who could talk about cakes all night! It also turned out that the guy hadn't stolen the database, he was just trying to guess by taste and improvising a bit.

After the whole mess, they both agreed to use your code, but asked you to do one last thing for them. They were quite impressed with how you managed to do the database comparison, so they asked you to do the same thing with their server's filesystem backups, so that neither bakery would have any technical problems in the future.

So your program should take two filesystem dumps:

`~$ ./compareFS --old snapshot1.txt --new snapshot2.txt`

They are both plain text files, unsorted, and each contains a file path on each like, like this:

```
/etc/stove/config.xml
/Users/baker/recipes/database.xml
/Users/baker/recipes/database_version3.yaml
/var/log/orders.log
/Users/baker/pokemon.avi
```

Your tool should output something very similar to the previous code (but without the CHANGED case):

```
ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
```

There is one problem though — the files can be very large, so you can assume that they won't fit in RAM at the same time. There are two ways to get around this: either compress the file in memory somehow, or just read one of them and then avoid reading the other. **Side note:** This is actually a very popular interview question.
