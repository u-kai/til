Patterns for Stack Components The following set of patterns and antipatterns give ideas for designing stack components and evaluating existing components. It’s not a complete list of ways you should or shouldn’t build modules and libraries; rather, it’s a starting point for thinking about the subject. Pattern: Facade Module Also known as: wrapper module. A facade module creates a simplified interface to a resource from the stack tool language or the infrastructure platform. The module exposes a few parameters to the calling code (see Example 16-1). Example 16-1. Example code using a facade module use module: shopspinner-server
name: checkout-appserver
memory: 8GB The module uses these parameters to call the resource it wraps, and hardcodes values for other parameters needed by the resource (Example 16-2).
Example 16-2. Code for the example facade module declare module: shopspinner-server
virtual_machine:
name: ${name}
source_image: hardened-linux-base
memory: ${memory}
provision:
tool: servermaker
maker_server: maker.shopspinner.xyz
role: application_server
network:
vlan: application_zone_vlan This example module allows the caller to create a virtual server, specifying the name and the amount of memory for the server. Every server created using the module uses a source image, role, and networking defined by the module. Motivation A facade module simplifies and standardizes a common use case for an infrastructure resource. The stack code that uses a facade module should be simpler and easier to read. Improvements to the quality of the module code are rapidly available to all of the stacks that use it. Applicability Facade modules work best for simple use cases, usually involving a basic infrastructure resource. Consequences A facade module limits how you can use the underlying infrastructure resource. Doing this can be useful, simplifying options and standardizing around better and more secure implementations. But it limits flexibility, so won’t apply to every use case.
A module is an extra layer of code between the stack code and the code that directly specifies the infrastructure resources. This extra layer adds at least some overhead to maintaining, debugging, and improving code. It can also make it harder to understand the stack code. Implementation Implementing a facade module generally involves specifying an infrastructure resource with a number of hardcoded values, and a small number of values that are passed through from the code that uses the module. A declarative infrastructure language is appropriate for a facade module. Related patterns An obfuscation module is a facade module that doesn’t hide much, adding complexity without adding much value. A bundle module (“Pattern: Bundle Module”) declares multiple related infrastructure resources, so is like a facade module with more parts. Antipattern: Obfuscation Module An obfuscation module wraps the code for an infrastructure element defined by the stack language or infrastructure platform, but does not simplify it or add any particular value. In the worst cases, the module complicates the code. See Example 16-3. Example 16-3. Example code using an obfuscation module use module: any_server
server_name: checkout-appserver
ram: 8GB
source_image: base_linux_image
provisioning_tool: servermaker
server_role: application_server
vlan: application_zone_vlan
The module itself passes the parameters directly to the stack management tool’s code, as shown in Example 16-4. Example 16-4. Code for the example obfuscation module declare module: any_server
virtual_machine:
name: ${server_name}
source_image: ${origin_server_image}
memory: ${ram}
provision:
tool: ${provisioning_tool}
role: ${server_role}
network:
vlan: ${server_vlan} Motivation An obfuscation module may be a facade module (see “Pattern: Facade Module”) gone wrong. Sometimes people write this kind of module aiming to follow the DRY principle (see “Avoid duplication”). They see that code that defines a common infrastructure element, such as a virtual server, load balancer, or security group, is used in multiple places in the codebase. So they create a module that declares that element type once and use that everywhere. But because the elements are being used differently in different parts of the code, they need to expose a large number of parameters in their module. Other people create obfuscation modules in a quest to design their own language for referring to infrastructure elements, “improving” the one provided by their stack tool. Applicability Nobody intentionally writes an obfuscation module. You may debate whether a given module obfuscates or is a facade, and that debate is useful. You should consider whether a module adds real value and, if not, then refactor it into code that uses the stack language directly.
Consequences Writing, using, and maintaining module code rather than directly using the constructs provided by your stack tool adds overhead. It adds more code to maintain, cognitive overhead to learn, and extra moving parts in your build and delivery process. A component should add enough value to make the overhead worthwhile. Implementation If a module neither simplifies the resources it defines nor adds value over the underlying stack language code, consider replacing usages by directly using the stack language. Related patterns An obfuscation module is similar to a facade module (see “Pattern: Facade Module”), but doesn’t noticeably simplify the underlying code.
